export class User {
  created_at: string;
  updated_at: string;
  username: string;
  first_name: string;
  last_name: string;
  password?: string;
  user_type: 'ADMIN' | 'USER';
  roles: string[];
  clusterIdList: string[];
  user_is_active: boolean;
  is_verified: boolean;
  forgot_password_token: string;
  phone: string;
}
