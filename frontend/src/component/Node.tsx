import PodBox from './PodBox';
import './node.css';

type Props = {
  nodeName: string;
  podNum: number;
}

export default function Node(props: Props) {
  const el = Array.from(Array(props.podNum).keys()).map((i: number) =>
    <PodBox key={i}/>
  )
  return (
      <fieldset className="node">
        <legend className="node-name-text">
          {props.nodeName}
        </legend>
        <div className="pods">
          {el}
        </div>
      </fieldset>
  )
}
