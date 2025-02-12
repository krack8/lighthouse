import { Directive, Input, TemplateRef, ViewContainerRef } from '@angular/core';
import { PermissionService } from '@core-ui/services';

/**
 * @whatItDoes Conditionally includes an HTML element if current user has any
 * of the authorities passed as the `expression`.
 * @howToUse
 * ```
 * <some-element *hasAnyAuthority="'CREATE_NAMESPACE'">...</some-element>
 * <some-element *hasAnyAuthority="['VIEW_NAMESPACE', 'CREATE_NAMESPACE', 'UPDATE_NAMESPACE']">...</some-element>
 * ```
 */
@Directive({
  standalone: true,
  selector: '[hasAnyAuthority]'
})
export class HasAnyAuthorityDirective {
  private permissions: string | string[] = [];

  constructor(
    private permissionSvc: PermissionService,
    private templateRef: TemplateRef<any>,
    private viewContainerRef: ViewContainerRef
  ) {}

  @Input()
  set hasAnyAuthority(value: string | string[]) {
    this.permissions = value;
    this.updateView();
  }

  private updateView(): void {
    const hasAnyAuthority = this.permissionSvc.hasAuthorities(this.permissions);
    this.viewContainerRef.clear();
    if (hasAnyAuthority) {
      this.viewContainerRef.createEmbeddedView(this.templateRef);
    }
  }
}
