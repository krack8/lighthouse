export interface IEndpoint {
  route: string;
  method: string;
}

export interface IRole {
  id: string;
  name: string;
  description: string;
  permissions: IPermission[];
  status;
  created_at: string;
  updated_at: string;
  created_by: string;
  updated_by: string;
}

export interface IPermission {
  id: string;
  name: string;
  description?: string;
  endpoint_list: IEndpoint[];
  category?: string;
  status: string;
  created_at: string;
  updated_at: string;
  created_by: string;
  updated_by: string;
}

export interface IFormPermission extends IPermission {
  checked: boolean;
}
