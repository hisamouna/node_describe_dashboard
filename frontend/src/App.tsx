import { useState, useEffect } from 'react';
import Node from './component/Node';
import Api, { nodeDescribeInf } from './api/api'
import './app.css';

export default function App() {
  var ndInf : nodeDescribeInf | undefined
  const [data, setData] = useState(ndInf)

  useEffect(() => {
    Api.nodeDescribe().then(
      ( response ) => {
        setData(response.data)
      }
    ).catch(
      err => console.log(`Error: ${err}`)
    )
  },[])

  return (
    <div className="dashboard">
      {
        data &&
          data.nodes.map((v,i) =>
            <Node podNum={v.pods.length} nodeName={v.name} key={i}/>
          )
      }
    </div>
  )
}
