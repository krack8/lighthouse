export const MC_GET_CLUSTERS = '/v1/clusters';
export const MC_GET_CLUSTER = '/v1/clusters/{0}';

export const MC_GET_MANUAL_CLUSTER = '/v1/cluster-step/manual';
export const MC_GET_CLUSTER_CREATION_STEP_STATUS = '/v1/cluster-step';
export const MC_GET_CLUSTER_FULL_LOG = '/v1/cluster-log';
// export const MC_GET_AVAILABLE_K8S_VERSION_LIST = '/v1/k8s-version';
export const MC_GET_AVAILABLE_K8S_VERSION_LIST = '/v1/eks-k8s-version-list';
export const MC_GET_AVAILABLE_AWS_REGION_LIST = '/v1/eks-region-list';
export const MC_GET_AVAILABLE_NODE_TYPE_LIST = '/v1/gke-node-group-data';
export const MC_GET_AWS_AVAILABLE_NODE_TYPE_LIST = '/v1/eks-node-machine-type';
export const MC_GET_AVAILABLE_K8S_VERSION_LIST_FOR_GCP = '/v1/gke-master-version-list';
export const MC_DELETE_CLUSTER = '/v1/multi-cluster/';
export const MC_VERIFY_DNS_SETTINGS = '/v1/cluster-dns-check/';
export const MC_FORCE_DELETE = '/v1/cluster/force-delete/';

export const MC_AWS_CLUSTER = '/v1/aws-cluster';
export const MC_GCP_CLUSTER = '/v1/gcp-cluster';
export const MC_DIGITAL_OCEAN_CLUSTER = '/v1/digital-ocean-cluster';
export const MC_ONBOARD_EXISTING_CLUSTER = '/v1/manual-cluster';

export const UPATE_CLUSTER = 'api/kube-cluster/update';
export const GET_CLUSTER_BY_ID = 'api/kube-cluster/v2/get/{0}';
export const GET_CLUSTER_RELEASE_BY_ID = 'api/v1/kube-cluster/cluster-release';
export const ACTIVATE_CLUSTER_UPGRADE = 'api/v1/kube-cluster/strategy/active/';
export const UPDATE_CLUSTER_UPGRADE_STRAETGY = 'api/v1/kube-cluster/strategy/';
