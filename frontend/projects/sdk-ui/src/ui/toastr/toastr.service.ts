// Imports from @angular
import { Injectable, ComponentFactoryResolver, ApplicationRef, Injector } from '@angular/core';
import { ComponentPortal, DomPortalHost } from '@angular/cdk/portal';
import { ToastrComponent } from './toastr.component';
import { Toastr } from './toastr.interface';
import { toasts } from './toastr.data';

@Injectable()
export class ToastrService {
  private toasterPortal: ComponentPortal<ToastrComponent>;
  private portalHost: DomPortalHost;

  constructor(
    private componentFactoryResolver: ComponentFactoryResolver,
    private appRef: ApplicationRef,
    private injector: Injector
  ) {
    this.toasterPortal = new ComponentPortal(ToastrComponent);

    this.portalHost = new DomPortalHost(document.body, this.componentFactoryResolver, this.appRef, this.injector);
  }

  success(message: string, title?: string): void {
    const toast: Toastr = {
      icon: 'icon-success',
      title: title || 'Success',
      message: message || ''
    };
    this.open(toast);
  }

  error(message: string, title?: string): void {
    const toast: Toastr = {
      icon: 'icon-error',
      title: title || 'Failed',
      message: message || ''
    };
    this.open(toast);
  }
  warn(message: string, title?: string): void {
    const toast: Toastr = {
      icon: 'icon-warn',
      title: title || '',
      message: message || ''
    };
    this.open(toast);
  }
  notification(message: string, title?: string): void {
    const toast: Toastr = {
      icon: 'icon-notification',
      title: title || 'Notification',
      message: message || ''
    };
    this.open(toast);
  }

  open(toast: Toastr) {
    if (!this.portalHost.hasAttached()) {
      this.portalHost.attach(this.toasterPortal);
    }
    toasts.push(toast);
    this.scheduleDismiss(toast.dismissAfter);
    return toasts.length - 1;
  }

  dismiss(index: number) {
    if (toasts[index]) {
      clearTimeout(toasts[index]['timeoutRef']);
      toasts[index]['dismiss'] = true;
    }
    toasts.splice(index, 1);
    if (toasts.length === 0) {
      this.portalHost.detach();
    }
  }

  private scheduleDismiss(dismissAfter: number = 10000) {
    toasts[toasts.length - 1]['timeoutRef'] = setTimeout(() => {
      toasts[0]['dismiss'] = true;
      setTimeout(() => toasts.shift(), 480);
    }, dismissAfter);
  }
}
