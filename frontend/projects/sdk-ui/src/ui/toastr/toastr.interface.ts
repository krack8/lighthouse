export interface Toastr {
  icon?: string;
  iconBg?: string;
  title: string;
  message?: string;
  dismissAfter?: number;
  timeoutRef?: any;
  dismiss?: boolean;
}
