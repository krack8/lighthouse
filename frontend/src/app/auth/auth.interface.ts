export interface LoginPayload {
  username: string;
  password: string;
}

export interface ResetPasswordPayload {
  token: string;
  newPassword: string;
}
