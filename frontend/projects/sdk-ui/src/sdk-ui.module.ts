import { ModuleWithProviders, NgModule } from '@angular/core';
import { MAT_FORM_FIELD_DEFAULT_OPTIONS, MatFormFieldDefaultOptions } from '@angular/material/form-field';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { CreateDropDownLink, SidenavLink, ToolbarUserMenuItem } from '@sdk-ui/interfaces';
import { LayoutModule } from '@sdk-ui/layout';
import { SdkConfigService } from '@sdk-ui/services';

@NgModule({
  imports: [LayoutModule, MatSnackBarModule],
  providers: [
    {
      provide: MAT_FORM_FIELD_DEFAULT_OPTIONS,
      useValue: {
        appearance: 'fill'
      } as MatFormFieldDefaultOptions
    }
  ]
})
export class SdkModule {
  static appConfig(options: {
    navigation: SidenavLink[];
    creationNavigation: CreateDropDownLink[];
    toolbarUserDropdown: ToolbarUserMenuItem[];
  }): ModuleWithProviders<SdkModule> {
    SdkConfigService.injectConfig({
      navigation: options.navigation,
      creationNavigation: options.creationNavigation,
      toolbarUserDropdown: options.toolbarUserDropdown
    });
    return {
      ngModule: SdkModule,
      providers: []
    };
  }
}
