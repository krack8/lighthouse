export type IDialogStaticConfig = {
  type: DialogStateType;

  // Static Config Props: if any static props is provided that props is highly prioritize
  message: string;
  note?: string;
  noteStyle?: { [k: string]: any };
  successBtnText: string;
  cancelBtnText: string;
  cancelBtnValue?: any;
  icon: string;
};

export type DialogStateType = 'delete' | 'info';
