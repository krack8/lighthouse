// task: Task = {
//   name: 'Indeterminate',
//   completed: false,
//   color: 'primary',
//   subtasks: [
//     {name: 'Primary', completed: false, color: 'primary'},
//     {name: 'Accent', completed: false, color: 'accent'},
//     {name: 'Warn', completed: false, color: 'warn'},
//   ],
// };

export interface IRole {
  id: string;
  description: string;
  roleName: string;
  listOfPermission: IPermissionListObject;
  createdDate: string;
  createdBy: string;
  updatedDate: string;
}

export interface IPermission {
  id: string;
  name: string;
  description?: string;
}

export interface IPermissionListObject {
  [key: string]: IPermission[];
}

export interface IFormPermission extends IPermission {
  checked: boolean;
}

export interface IFormCategoryPermission {
  name: string;
  label?: string;
  completed: boolean;
  permissions: IFormPermission[];
}
