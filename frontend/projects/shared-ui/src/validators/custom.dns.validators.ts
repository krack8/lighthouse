import { AbstractControl, ValidationErrors } from '@angular/forms';

export class DnsValidator {
  static domain(control: AbstractControl): ValidationErrors | null {
    const domainFormat =
      /^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,7})$/;
    if (control.value && !(control.value as string).match(domainFormat)) {
      return { domain: true };
    }

    return null;
  }

  /**
   * @description http url validation
   * @uasge keep 'url' validation on last
   * ```ts
   * { controlName: ['', [Validation.required, DnsValidator.url]] }
   * ```
   * @returns {{url: string} | null} {url: string} | null
   */
  static url(control: AbstractControl): ValidationErrors | null {
    const format = /^((http|https)?:\/\/)?([a-zA-Z0-9-]+\.){1,}[a-zA-Z]{2,}$/;
    if (control.value && !(control.value as string).match(format)) {
      return { url: 'Url is not valid' };
    }
    return null;
  }

  /**
   * @descriptionws url validation
   * @uasge keep 'wsUrl' validation on last
   * ```ts
   * { controlName: ['', [Validation.required, DnsValidator.wsUrl]] }
   * ```
   * @returns {{url: string} | null} {url: string} | null
   */
  static wsUrl(control: AbstractControl): ValidationErrors | null {
    const format = /^((ws|wss):\/\/)?([a-zA-Z0-9-]+\.){1,}[a-zA-Z]{2,}$/;
    if (control.value && !(control.value as string).match(format)) {
      return { url: 'WS url is not valid' };
    }
    return null;
  }

  static ipV4(control: AbstractControl): ValidationErrors | null {
    const ipPattern =
      /^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/;
    if (!(control.value as string).match(ipPattern)) {
      return { ipV4: true };
    }

    return null;
  }

  static ipV6(control: AbstractControl): ValidationErrors | null {
    const ipPattern =
      /(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))/;
    if (!(control.value as string).match(ipPattern)) {
      return { ipV6: true };
    }

    return null;
  }

  static ip(control: AbstractControl): ValidationErrors | null {
    const ipv4 = /(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/;

    const ipv6 = /((([0-9a-fA-F]){1,4})\:){7}([0-9a-fA-F]){1,4}/;
    if ((control.value as string).match(ipv4) || (control.value as string).match(ipv6)) return null;

    return { invalidIp: true };
  }

  /**
   * @descriptionws host validation. Example: console.xyz.co
   * @uasge keep 'host' validation on last
   * ```ts
   * { controlName: ['', [Validation.required, DnsValidator.host]] }
   * ```
   * @returns {{host: boolean} | null} {host: true} | null
   */
  static host(control: AbstractControl): ValidationErrors | null {
    const hostPattern =
      /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\-]*[A-Za-z0-9])$/;
    if (control.value && !(control.value as string).match(hostPattern)) {
      return { host: true };
    }

    return null;
  }

  static port(control: AbstractControl): ValidationErrors | null {
    const portPattern = /^()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5])$/;
    if (control.value && !(control.value as string).match(portPattern)) {
      return { port: true };
    }

    return null;
  }
}
