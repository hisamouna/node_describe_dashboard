import http from './http-common';

export interface nodeDescribeInf {
  nodes: node[]
}

export interface node {
  name: string
  pods: pod[]
}

interface pod {
  namespace: string
  name: string
}

const nodeDescribe = () => {
  return http.server.get<nodeDescribeInf>("/node/describe");
}

const Api = {
  nodeDescribe,
}

export default Api
