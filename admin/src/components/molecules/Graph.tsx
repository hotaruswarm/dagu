import React from 'react';
import { Node, NodeStatus } from '../../models';
import { Step } from '../../models';
import Mermaid from '../atoms/Mermaid';

type onClickNode = (name: string) => void;

export type FlowchartType = 'TD' | 'LR';

type Props = {
  type: 'status' | 'config';
  flowchart?: FlowchartType;
  steps?: Step[] | Node[];
  onClickNode?: onClickNode;
};

declare global {
  interface Window {
    onClickMermaidNode: onClickNode;
  }
}

function Graph({
  steps,
  flowchart = 'TD',
  type = 'status',
  onClickNode,
}: Props) {
  const mermaidStyle = {
    display: 'flex',
    alignItems: 'flex-center',
    justifyContent: 'flex-start',
    width: flowchart == 'LR' && steps ? steps.length * 240 + 'px' : '100%',
    minWidth: '100%',
    minHeight: '200px',
    padding: '2em',
    borderRadius: '0.5em',
    backgroundSize: '20px 20px',
  };
  const graph = React.useMemo(() => {
    if (!steps) {
      return '';
    }
    const dat = flowchart == 'TD' ? ['flowchart TD;'] : ['flowchart LR;'];
    if (onClickNode) {
      window.onClickMermaidNode = onClickNode;
    }
    const addNodeFn = (step: Step, status: NodeStatus) => {
      const id = step.Name.replace(/\s/g, '_');
      const c = graphStatusMap[status] || '';
      dat.push(`${id}(${step.Name})${c};`);
      if (step.Depends) {
        step.Depends.forEach((d) => {
          const depId = d.replace(/\s/g, '_');
          dat.push(`${depId} -.-> ${id};`);
        });
      }
      if (onClickNode) {
        dat.push(`click ${id} onClickMermaidNode`);
      }
    };
    if (type == 'status') {
      (steps as Node[]).forEach((s) => addNodeFn(s.Step, s.Status));
    } else {
      (steps as Step[]).forEach((s) => addNodeFn(s, NodeStatus.None));
    }
    dat.push(
      'linkStyle default stroke:#ddeeff,stroke-width:2px,fill:none,color:#404040'
    );
    dat.push('classDef none fill:#bbbbff,stroke-width:0px,color:#000');
    dat.push('classDef running fill:#33ff33,stroke-width:0px,color:#000');
    dat.push('classDef error fill:#ee0000,stroke-width:0px,color:#000');
    dat.push('classDef cancel fill:#ffbbaa,stroke-width:0px,color:#000');
    dat.push('classDef done fill:#00bb00,stroke-width:0px,color:#000');
    dat.push('classDef skipped fill:#dfdfdf,stroke-width:0px,color:#000');
    return dat.join('\n');
  }, [steps, onClickNode, flowchart]);
  return <Mermaid style={mermaidStyle} def={graph} />;
}

export default Graph;

const graphStatusMap = {
  [NodeStatus.None]: ':::none',
  [NodeStatus.Running]: ':::running',
  [NodeStatus.Error]: ':::error',
  [NodeStatus.Cancel]: ':::cancel',
  [NodeStatus.Success]: ':::done',
  [NodeStatus.Skipped]: ':::skipped',
};
