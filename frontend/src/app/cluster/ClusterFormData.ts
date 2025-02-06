// #1 AWS
import { NodeType } from './cluster.interface';

export const awsRegionList = [
  {
    title: 'US East (Ohio)',
    name: 'us-east-2'
  },
  {
    title: 'US East (N. Virginia)',
    name: 'us-east-1'
  },
  {
    title: 'US West (N. California)',
    name: 'us-west-1'
  },
  {
    title: 'US West (Oregon)',
    name: 'us-west-2'
  },
  {
    title: 'Africa (Cape Town)',
    name: 'af-south-1'
  },
  {
    title: 'Asia Pacific (Hong Kong)',
    name: 'ap-east-1'
  },
  {
    title: 'Asia Pacific (Jakarta)',
    name: 'ap-southeast-3'
  },
  {
    title: 'Asia Pacific (Mumbai)',
    name: 'ap-south-1'
  },
  {
    title: 'Asia Pacific (Osaka)',
    name: 'ap-northeast-3'
  },
  {
    title: 'Asia Pacific (Seoul)',
    name: 'ap-northeast-2'
  },
  {
    title: 'Asia Pacific (Singapore)',
    name: 'ap-southeast-1'
  },
  {
    title: 'Asia Pacific (Sydney)',
    name: 'ap-southeast-2'
  },
  {
    title: 'Asia Pacific (Tokyo)',
    name: 'ap-northeast-1'
  },
  {
    title: 'Canada (Central)',
    name: 'ca-central-1'
  },
  {
    title: 'Europe (Frankfurt)',
    name: 'eu-central-1'
  },
  {
    title: 'Europe (Ireland)',
    name: 'eu-west-1'
  },
  {
    title: 'Europe (London)',
    name: 'eu-west-2'
  },
  {
    title: 'Europe (Milan)',
    name: 'eu-south-1'
  },
  {
    title: 'Europe (Paris)',
    name: 'eu-west-3'
  },
  {
    title: 'Europe (Stockholm)',
    name: 'eu-north-1'
  },
  {
    title: 'Middle East (Bahrain)',
    name: 'me-south-1'
  },
  {
    title: 'South America (São Paulo)',
    name: 'sa-east-1'
  },
  {
    title: 'AWS GovCloud (US-East)',
    name: 'us-gov-east-1'
  },
  {
    title: 'AWS GovCloud (US-West)',
    name: 'us-gov-west-1'
  }
];

// Small Instance size: CPU=1, MEMORY=2
// Large instance use for Cluster master node. Size: CPU=2, MEMORY=8
export const awsNodeTypeList: NodeType[] = [
  // T2
  { name: 't2.medium', vcpu: 2, memory: 4 },
  { name: 't2.large', vcpu: 2, memory: 8 },
  { name: 't2.xlarge', vcpu: 4, memory: 16 },
  { name: 't2.2xlarge', vcpu: 8, memory: 32 },
  // T3
  { name: 't3.small', vcpu: 2, memory: 2 },
  { name: 't3.medium', vcpu: 2, memory: 4 },
  { name: 't3.large', vcpu: 2, memory: 8 },
  { name: 't3.xlarge', vcpu: 4, memory: 16 },
  { name: 't3.2xlarge', vcpu: 8, memory: 32 },
  // T3a
  { name: 't3a.small', vcpu: 2, memory: 2 },
  { name: 't3a.medium', vcpu: 2, memory: 4 },
  { name: 't3a.large', vcpu: 2, memory: 8 },
  { name: 't3a.xlarge', vcpu: 4, memory: 16 },
  { name: 't3a.2xlarge', vcpu: 8, memory: 32 },
  // T4G
  { name: 't3g.small', vcpu: 2, memory: 2 },
  { name: 't4g.medium', vcpu: 2, memory: 4 },
  { name: 't4g.large', vcpu: 2, memory: 8 },
  { name: 't4g.xlarge', vcpu: 4, memory: 16 },
  { name: 't4g.2xlarge', vcpu: 8, memory: 32 },
  // MAC
  { name: 'mac1.metal', vcpu: 12, memory: 32 },
  { name: 'mac2.metal', vcpu: 8, memory: 16 }
];

// #2 GCP

export const gcpNodeTypeList: NodeType[] = [
  // e2
  { name: 'e2-small', vcpu: 2, memory: 2 },
  { name: 'e2-medium', vcpu: 2, memory: 4 },

  { name: 'e2-standard-2', vcpu: 2, memory: 8 },
  { name: 'e2-standard-4', vcpu: 4, memory: 16 },
  { name: 'e2-standard-8', vcpu: 8, memory: 32 },
  { name: 'e2-standard-16', vcpu: 16, memory: 64 },
  { name: 'e2-standard-32', vcpu: 32, memory: 120 },

  { name: 'e2-highmem-2', vcpu: 2, memory: 16 },
  { name: 'e2-highmem-4', vcpu: 4, memory: 32 },
  { name: 'e2-highmem-8', vcpu: 8, memory: 64 },
  { name: 'e2-highmem-16', vcpu: 16, memory: 128 },

  { name: 'e2-highcpu-2', vcpu: 2, memory: 2 },
  { name: 'e2-highcpu-4', vcpu: 4, memory: 4 },
  { name: 'e2-highcpu-8', vcpu: 8, memory: 8 },
  { name: 'e2-highcpu-16', vcpu: 16, memory: 16 },
  { name: 'e2-highcpu-32', vcpu: 32, memory: 32 },
  // N1
  { name: 'n1-standard-2', vcpu: 2, memory: 7.5 },
  { name: 'n1-standard-4', vcpu: 4, memory: 15 },
  { name: 'n1-standard-8', vcpu: 8, memory: 30 },
  { name: 'n1-standard-16', vcpu: 16, memory: 60 },
  { name: 'n1-standard-32', vcpu: 32, memory: 120 },
  { name: 'n1-standard-64', vcpu: 64, memory: 240 },
  { name: 'n1-standard-96', vcpu: 96, memory: 360 },

  { name: 'n1-highmem-2', vcpu: 2, memory: 13 },
  { name: 'n1-highmem-4', vcpu: 4, memory: 26 },
  { name: 'n1-highmem-8', vcpu: 8, memory: 52 },
  { name: 'n1-highmem-16', vcpu: 16, memory: 104 },
  { name: 'n1-highmem-32', vcpu: 32, memory: 208 },
  { name: 'n1-highmem-64', vcpu: 64, memory: 416 },
  { name: 'n1-highmem-96', vcpu: 96, memory: 624 },

  { name: 'n1-highcpu-4', vcpu: 4, memory: 3.6 },
  { name: 'n1-highcpu-8', vcpu: 8, memory: 7.2 },
  { name: 'n1-highcpu-16', vcpu: 16, memory: 14.4 },
  { name: 'n1-highcpu-32', vcpu: 32, memory: 28.8 },
  { name: 'n1-highcpu-64', vcpu: 64, memory: 57.6 },
  { name: 'n1-highcpu-96', vcpu: 96, memory: 86.4 },

  // N2
  { name: 'n2-standard-2', vcpu: 2, memory: 8 },
  { name: 'n2-standard-4', vcpu: 4, memory: 16 },
  { name: 'n2-standard-8', vcpu: 8, memory: 32 },
  { name: 'n2-standard-16', vcpu: 16, memory: 64 },
  { name: 'n2-standard-32', vcpu: 32, memory: 128 },
  { name: 'n2-standard-48', vcpu: 48, memory: 192 },
  { name: 'n2-standard-64', vcpu: 64, memory: 256 },
  { name: 'n2-standard-80', vcpu: 80, memory: 320 },
  { name: 'n2-standard-96', vcpu: 96, memory: 384 },
  { name: 'n2-standard-128', vcpu: 128, memory: 512 },

  { name: 'n2-highmem-2', vcpu: 2, memory: 16 },
  { name: 'n2-highmem-4', vcpu: 4, memory: 32 },
  { name: 'n2-highmem-8', vcpu: 8, memory: 64 },
  { name: 'n2-highmem-16', vcpu: 16, memory: 128 },
  { name: 'n2-highmem-32', vcpu: 32, memory: 256 },
  { name: 'n2-highmem-48', vcpu: 48, memory: 384 },
  { name: 'n2-highmem-64', vcpu: 64, memory: 512 },
  { name: 'n2-highmem-80', vcpu: 80, memory: 640 },
  { name: 'n2-highmem-96', vcpu: 96, memory: 768 },
  { name: 'n2-highmem-128', vcpu: 128, memory: 864 },

  { name: 'n2-highcpu-2', vcpu: 2, memory: 2 },
  { name: 'n2-highcpu-4', vcpu: 4, memory: 4 },
  { name: 'n2-highcpu-8', vcpu: 8, memory: 8 },
  { name: 'n2-highcpu-16', vcpu: 16, memory: 16 },
  { name: 'n2-highcpu-32', vcpu: 32, memory: 32 },
  { name: 'n2-highcpu-48', vcpu: 48, memory: 48 },
  { name: 'n2-highcpu-64', vcpu: 64, memory: 64 },
  { name: 'n2-highcpu-80', vcpu: 80, memory: 80 },
  { name: 'n2-highcpu-96', vcpu: 96, memory: 96 },
  { name: 'n2-highcpu-128', vcpu: 128, memory: 128 },

  // N2D
  { name: 'n2d-standard-2', vcpu: 2, memory: 8 },
  { name: 'n2d-standard-4', vcpu: 4, memory: 16 },
  { name: 'n2d-standard-8', vcpu: 8, memory: 32 },
  { name: 'n2d-standard-16', vcpu: 16, memory: 64 },
  { name: 'n2d-standard-32', vcpu: 32, memory: 128 },
  { name: 'n2d-standard-48', vcpu: 48, memory: 192 },
  { name: 'n2d-standard-64', vcpu: 64, memory: 256 },
  { name: 'n2d-standard-80', vcpu: 80, memory: 320 },
  { name: 'n2d-standard-96', vcpu: 96, memory: 384 },
  { name: 'n2d-standard-128', vcpu: 128, memory: 512 },
  { name: 'n2d-standard-224', vcpu: 224, memory: 896 },

  { name: 'n2d-highmem-2', vcpu: 2, memory: 16 },
  { name: 'n2d-highmem-4', vcpu: 4, memory: 32 },
  { name: 'n2d-highmem-8', vcpu: 8, memory: 64 },
  { name: 'n2d-highmem-16', vcpu: 16, memory: 128 },
  { name: 'n2d-highmem-32', vcpu: 32, memory: 256 },
  { name: 'n2d-highmem-48', vcpu: 48, memory: 384 },
  { name: 'n2d-highmem-64', vcpu: 64, memory: 512 },
  { name: 'n2d-highmem-80', vcpu: 80, memory: 640 },
  { name: 'n2d-highmem-96', vcpu: 96, memory: 768 },

  { name: 'n2d-highcpu-2', vcpu: 2, memory: 2 },
  { name: 'n2d-highcpu-4', vcpu: 4, memory: 4 },
  { name: 'n2d-highcpu-8', vcpu: 8, memory: 8 },
  { name: 'n2d-highcpu-16', vcpu: 16, memory: 16 },
  { name: 'n2d-highcpu-32', vcpu: 32, memory: 32 },
  { name: 'n2d-highcpu-48', vcpu: 48, memory: 48 },
  { name: 'n2d-highcpu-64', vcpu: 64, memory: 64 },
  { name: 'n2d-highcpu-80', vcpu: 80, memory: 80 },
  { name: 'n2d-highcpu-96', vcpu: 96, memory: 96 },
  { name: 'n2d-highcpu-128', vcpu: 128, memory: 128 },
  { name: 'n2d-highcpu-224', vcpu: 224, memory: 224 }
];

// Regions
export const gcpRegionList = [
  // AMERICAS
  {
    title: 'OREGON',
    name: 'us-west1'
  },
  {
    title: 'LOS ANGELES',
    name: 'us-west2'
  },
  {
    title: 'SALT LAKE CITY',
    name: 'us-west3'
  },
  {
    title: 'LAS VEGAS',
    name: 'us-west4'
  },
  {
    title: 'IOWA',
    name: 'us-central1'
  },
  {
    title: 'SOUTH CAROLINA',
    name: 'us-east1'
  },
  {
    title: 'N. VIRGINIA',
    name: 'us-east4'
  },
  {
    title: 'MONTRÉAL',
    name: 'northamerica-northeast1'
  },
  {
    title: 'SÃO PAULO',
    name: 'southamerica-east1'
  },
  // EUROPE
  {
    title: 'LONDON',
    name: 'europe-west2'
  },
  {
    title: 'BELGIUM',
    name: 'europe-west1'
  },
  {
    title: 'NETHERLANDS',
    name: 'europe-west4'
  },
  {
    title: 'ZURICH',
    name: 'europe-west6'
  },
  {
    title: 'FRANKFURT',
    name: 'europe-west3'
  },
  {
    title: 'FINLAND',
    name: 'europe-north1'
  },
  // ASIA PACIFIC
  {
    title: 'MUMBAI',
    name: 'asia-south1'
  },
  {
    title: 'SINGAPORE',
    name: 'asia-southeast1'
  },
  {
    title: 'JAKARTA',
    name: 'asia-southeast2'
  },
  {
    title: 'HONG KONG',
    name: 'asia-east2'
  },
  {
    title: 'TAIWAN',
    name: 'asia-east1'
  },
  {
    title: 'TOKYO',
    name: 'asia-northeast1'
  },
  {
    title: 'OSAKA',
    name: 'asia-northeast2'
  },
  {
    title: 'SYDNEY',
    name: 'australia-southeast1'
  },
  {
    title: 'SEOUL',
    name: 'asia-northeast3'
  }
];

export const gcpZoneList = [
  /* ---------------------------- North America --------------------- */
  {
    region: 'northamerica-northeast1',
    name: 'northamerica-northeast1-a'
  },
  {
    region: 'northamerica-northeast1',
    name: 'northamerica-northeast1-b'
  },
  {
    region: 'northamerica-northeast1',
    name: 'northamerica-northeast1-c'
  },
  {
    region: 'northamerica-northeast2',
    name: 'northamerica-northeast2-a'
  },
  {
    region: 'northamerica-northeast2',
    name: 'northamerica-northeast2-b'
  },
  {
    region: 'northamerica-northeast2',
    name: 'northamerica-northeast2-c'
  },
  { region: 'us-central1', name: 'us-central1-a' },
  { region: 'us-central1', name: 'us-central1-b' },
  { region: 'us-central1', name: 'us-central1-c' },
  { region: 'us-central1', name: 'us-central1-f' },
  { region: 'us-east1', name: 'us-east1-b' },
  { region: 'us-east1', name: 'us-east1-c' },
  { region: 'us-east1', name: 'us-east1-d' },
  { region: 'us-east4', name: 'us-east4-a' },
  { region: 'us-east4', name: 'us-east4-b' },
  { region: 'us-east4', name: 'us-east4-c' },
  { region: 'us-east5', name: 'us-east5-a' },
  { region: 'us-east5', name: 'us-east5-b' },
  { region: 'us-east5', name: 'us-east5-c' },
  { region: 'us-west1', name: 'us-west1-a' },
  { region: 'us-west1', name: 'us-west1-b' },
  { region: 'us-west1', name: 'us-west1-c' },
  { region: 'us-west2', name: 'us-west2-a' },
  { region: 'us-west2', name: 'us-west2-b' },
  { region: 'us-west2', name: 'us-west2-c' },
  { region: 'us-west3', name: 'us-west3-a' },
  { region: 'us-west3', name: 'us-west3-b' },
  { region: 'us-west3', name: 'us-west3-c' },
  { region: 'us-west4', name: 'us-west4-a' },
  { region: 'us-west4', name: 'us-west4-b' },
  { region: 'us-west4', name: 'us-west4-c' },
  { region: 'us-south1', name: 'us-south1-a' },
  { region: 'us-south1', name: 'us-south1-b' },
  { region: 'us-south1', name: 'us-south1-c' },

  /* ---------------------------- South America --------------------- */
  { region: 'southamerica-east1', name: 'southamerica-east1-a' },
  { region: 'southamerica-east1', name: 'southamerica-east1-b' },
  { region: 'southamerica-east1', name: 'southamerica-east1-c' },
  { region: 'southamerica-west1', name: 'southamerica-west1-a' },
  { region: 'southamerica-west1', name: 'southamerica-west1-b' },
  { region: 'southamerica-west1', name: 'southamerica-west1-c' },

  /* ---------------------------- Europe --------------------- */
  { region: 'europe-north1', name: 'europe-north1-a' },
  { region: 'europe-north1', name: 'europe-north1-b' },
  { region: 'europe-north1', name: 'europe-north1-c' },
  { region: 'europe-central2', name: 'europe-central2-a' },
  { region: 'europe-central2', name: 'europe-central2-b' },
  { region: 'europe-central2', name: 'europe-central2-c' },
  { region: 'europe-southwest1', name: 'europe-southwest1-a' },
  { region: 'europe-southwest1', name: 'europe-southwest1-b' },
  { region: 'europe-southwest1', name: 'europe-southwest1-c' },
  { region: 'europe-west1', name: 'europe-west1-a' },
  { region: 'europe-west1', name: 'europe-west1-b' },
  { region: 'europe-west1', name: 'europe-west1-c' },
  { region: 'europe-west2', name: 'europe-west2-a' },
  { region: 'europe-west2', name: 'europe-west2-b' },
  { region: 'europe-west2', name: 'europe-west2-c' },
  { region: 'europe-west3', name: 'europe-west3-a' },
  { region: 'europe-west3', name: 'europe-west3-b' },
  { region: 'europe-west3', name: 'europe-west3-c' },
  { region: 'europe-west4', name: 'europe-west4-a' },
  { region: 'europe-west4', name: 'europe-west4-b' },
  { region: 'europe-west4', name: 'europe-west4-c' },
  { region: 'europe-west6', name: 'europe-west6-a' },
  { region: 'europe-west6', name: 'europe-west6-b' },
  { region: 'europe-west6', name: 'europe-west6-c' },
  { region: 'europe-west8', name: 'europe-west8-a' },
  { region: 'europe-west8', name: 'europe-west8-b' },
  { region: 'europe-west8', name: 'europe-west8-c' },
  { region: 'europe-west9', name: 'europe-west9-a' },
  { region: 'europe-west9', name: 'europe-west9-b' },
  { region: 'europe-west9', name: 'europe-west9-c' },

  /* ---------------------------- Asia Pacific --------------------- */
  { region: 'asia-east1', name: 'asia-east1-a' },
  { region: 'asia-east1', name: 'asia-east1-b' },
  { region: 'asia-east1', name: 'asia-east1-c' },
  { region: 'asia-east2', name: 'asia-east2-a' },
  { region: 'asia-east2', name: 'asia-east2-b' },
  { region: 'asia-east2', name: 'asia-east2-c' },
  { region: 'asia-northeast1', name: 'asia-northeast1-a' },
  { region: 'asia-northeast1', name: 'asia-northeast1-b' },
  { region: 'asia-northeast1', name: 'asia-northeast1-c' },
  { region: 'asia-northeast2', name: 'asia-northeast2-a' },
  { region: 'asia-northeast2', name: 'asia-northeast2-b' },
  { region: 'asia-northeast2', name: 'asia-northeast2-c' },
  { region: 'asia-northeast3', name: 'asia-northeast3-a' },
  { region: 'asia-northeast3', name: 'asia-northeast3-b' },
  { region: 'asia-northeast3', name: 'asia-northeast3-c' },
  { region: 'asia-south1', name: 'asia-south1-a' },
  { region: 'asia-south1', name: 'asia-south1-b' },
  { region: 'asia-south1', name: 'asia-south1-c' },
  { region: 'asia-south2', name: 'asia-south2-a' },
  { region: 'asia-south2', name: 'asia-south2-b' },
  { region: 'asia-south2', name: 'asia-south2-c' },
  { region: 'asia-southeast1', name: 'asia-southeast1-a' },
  { region: 'asia-southeast1', name: 'asia-southeast1-b' },
  { region: 'asia-southeast1', name: 'asia-southeast1-c' },
  { region: 'asia-southeast2', name: 'asia-southeast2-a' },
  { region: 'asia-southeast2', name: 'asia-southeast2-b' },
  { region: 'asia-southeast2', name: 'asia-southeast2-c' },
  { region: 'australia-southeast1', name: 'australia-southeast1-a' },
  { region: 'australia-southeast1', name: 'australia-southeast1-b' },
  { region: 'australia-southeast1', name: 'australia-southeast1-c' },
  { region: 'australia-southeast2', name: 'australia-southeast2-a' },
  { region: 'australia-southeast2', name: 'australia-southeast2-b' },
  { region: 'australia-southeast2', name: 'australia-southeast2-c' },

  /* ---------------------------- Middle East --------------------- */
  { region: 'me-west1', name: 'me-west1-a' },
  { region: 'me-west1', name: 'me-west1-b' },
  { region: 'me-west1', name: 'me-west1-c' }
];

// gcp k8s version
export const gcpK8sVersionList = [
  '1.27.3-gke.100',
  '1.27.2-gke.2100',
  '1.27.2-gke.1200',
  '1.26.5-gke.2700',
  '1.26.5-gke.2100',
  '1.26.5-gke.1400',
  '1.26.5-gke.1200',
  '1.25.10-gke.2700',
  '1.25.10-gke.2100',
  '1.25.10-gke.1400',
  '1.25.10-gke.1200',
  '1.25.9-gke.2300',
  '1.25.8-gke.1000',
  '1.24.14-gke.2700',
  '1.24.14-gke.2100',
  '1.24.14-gke.1400',
  '1.24.14-gke.1200',
  '1.24.13-gke.2500',
  '1.23.17-gke.8400',
  '1.23.17-gke.7700',
  '1.23.17-gke.7000',
  '1.23.17-gke.6800',
  '1.22.17-gke.14100',
  '1.22.17-gke.12700',
  '1.21.14-gke.18800'
];

// #3 Digital Ocean
export const digitalOceanRegionList = [
  'nyc1',
  'nyc2',
  'nyc3',
  'ams2',
  'ams3',
  'sfo1',
  'sfo2',
  'sfo3',
  'sgp1',
  'lon1',
  'fra1',
  'tor1',
  'blr1'
];
