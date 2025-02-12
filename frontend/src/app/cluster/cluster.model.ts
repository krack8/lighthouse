export type ClusterType = 'MASTER' | 'AGENT';

export interface ICluster {
  id: string;
  name: string;
  cluster_type: ClusterType;
  master_cluster_id: string;
  resource_namespace: string;
  is_active: boolean;
  status: string;
  created_at: string;
  updated_at: string;
  created_by: string;
  updated_by: string;
}
