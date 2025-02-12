import { Toastr } from '@sdk-ui/ui';

export class Utils {
  static formatString(text: string, ...args: string[]) {
    let a = text;

    for (const k in args) {
      a = a.replace('{' + k + '}', args[k]);
    }
    return a;
  }

  /**
   * @deprecated Use toastr service notification method.
   */
  static getDefaultConfig(toastr: Toastr = { icon: '', iconBg: '', title: '', message: '' }): Toastr {
    const { icon, iconBg, title, message } = toastr;
    return {
      icon: icon || 'icon-notification',
      iconBg: iconBg || 'text-success',
      title: title || 'Action Completed',
      message: message || ''
    };
  }
  /**
   * @deprecated Use toastr service success method.
   */
  static getCommonSuccessConfig(toastr: Toastr = { icon: '', iconBg: '', title: '', message: '' }): Toastr {
    const { icon, iconBg, title, message } = toastr;
    return {
      icon: icon || 'icon-success',
      iconBg: iconBg || 'text-success',
      title: title || 'Action Completed',
      message: message || ''
    };
  }
  /**
   * @deprecated Use toastr service warn method.
   */
  static getCommonWarningConfig(toastr: Toastr = { icon: '', iconBg: '', title: '', message: '' }): Toastr {
    const { icon, iconBg, title, message } = toastr;
    return {
      icon: icon || 'icon-warn',
      iconBg: iconBg || 'text-warn',
      title: title || 'Action Completed',
      message: message || ''
    };
  }
  /**
   * @deprecated Use toastr service error method.
   */
  static getCommonErrorConfig(toastr: Toastr = { icon: '', iconBg: '', title: '', message: '' }): Toastr {
    const { icon, iconBg, title, message } = toastr;
    return {
      icon: icon || 'icon-error',
      iconBg: iconBg || 'text-error',
      title: title || 'Oops!',
      message: message || 'Something went wrong!'
    };
  }

  static processLog(log: string, ignoreWarring: boolean = true) {
    const match = /\r|\n/.exec(log);
    if (match) {
      log = log.replace(/\n/g, '<br>');
      log = log.replace(/\r/g, '&emsp;');
    }
    const removeUnnecessaryPrefixRegex = /\[([0-9]?[0-9])m/g;
    log = log.replace(removeUnnecessaryPrefixRegex, '');
    // eslint-disable-next-line no-control-regex
    log = log.replace(/[\x00-\x09\x0b-\x1F]/g, ' ');
    // Prevent If warning
    if (ignoreWarring && log.match(/warning:|Warning:/g) != null) {
      return '';
    }
    log = log.replace(/INFO/g, '<span class="text-info">INFO</span>');
    log = log.replace(/BUILD SUCCESSFUL/g, '<span class="text-success">BUILD SUCCESSFUL</span>');
    log = log.replace(/BUILD SUCCESS/g, '<span class="text-success">BUILD SUCCESS</span>');
    log = log.replace(/SUCCESSFUL/g, '<span class="text-success">SUCCESSFUL</span>');
    log = log.replace(/SUCCESS/g, '<span class="text-success">SUCCESS</span>');
    log = log.replace(/WARN/g, '<span class="text-warn"> WARN </span>');
    log = log.replace(/WARNING/g, '<span class="text-warn">WARNING </span>');
    log = log.replace(/ERROR/g, '<span class="text-error">ERROR</span>');
    log = log.replace(/FAILED/g, '<span class="text-error">FAILED</span>');
    return log;
  }
}

// Checking Value null value
export const isNil = (value: any): value is null | undefined => {
  return value === null || typeof value === 'undefined';
};

export const isObject = (value: any): boolean => {
  return value && value.constructor === Object;
};

export const isBlank = (value: any): boolean => {
  return isNil(value) || (isObject(value) && Object.keys(value).length === 0) || value.toString().trim() === '';
};

export const isPresent = (value: any): boolean => {
  return !isBlank(value);
};
