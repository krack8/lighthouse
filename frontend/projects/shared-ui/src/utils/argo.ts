export type SyncStatusCode = 'Unknown' | 'Synced' | 'OutOfSync';

export const SyncStatuses: { [key: string]: SyncStatusCode } = {
  Unknown: 'Unknown',
  Synced: 'Synced',
  OutOfSync: 'OutOfSync'
};

export type HealthStatusCode = 'Unknown' | 'Progressing' | 'Healthy' | 'Suspended' | 'Degraded' | 'Missing';

export const HealthStatuses: { [key: string]: HealthStatusCode } = {
  Unknown: 'Unknown',
  Progressing: 'Progressing',
  Healthy: 'Healthy',
  Suspended: 'Suspended',
  Degraded: 'Degraded',
  Missing: 'Missing'
};

export interface HealthStatus {
  status: HealthStatusCode;
  message: string;
}

export namespace CD {
  export interface ResourceStatus {
    group: string;
    version: string;
    kind: string;
    namespace: string;
    name: string;
    message: string;
    syncPhase: string;
    status: SyncStatusCode;
    hookPhase: HealthStatusCode;
  }

  export interface ResourceRef {
    uid: string;
    kind: string;
    namespace: string;
    name: string;
    version: string;
    group: string;
  }

  export interface InfoItem {
    name: string;
    value: string;
  }
  export interface ResourceNode extends ResourceRef {
    parentRefs: ResourceRef[];
    info: InfoItem[];
    images?: string[];
    resourceVersion: string;
    health?: HealthStatus;
    createdAt?: any;
  }
}
