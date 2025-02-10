import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { RouterModule } from '@angular/router';
import { AppComponent } from './app.component';
import { appRoutes } from './app.routes';
import { NgxWebstorageModule } from 'ngx-webstorage';
import { InitializerModule } from '@core-ui/initializer.module';
import { SdkModule } from '@sdk-ui/sdk-ui.module';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { PermissionService, RequesterService } from '@core-ui/services';
import { LayoutService, StyleService } from '@sdk-ui/services';
import { APP_ENV } from '@core-ui/constants';
import { AuthInterceptor, ErrorsInterceptor } from '@core-ui/interceptors';
import { environment } from '../environments/environment';
import { RoleGuardService } from '@core-ui/guards';
import { SIDENAV_LIST } from '../data/navigation';
import { USER_DROPDOWN_LIST } from '../data/toolbar';
import { ToastrModule } from '@sdk-ui/ui';

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    ToastrModule,
    InitializerModule,
    NgxWebstorageModule.forRoot(),
    SdkModule.appConfig({ navigation: SIDENAV_LIST, creationNavigation: [], toolbarUserDropdown: USER_DROPDOWN_LIST }),
    RouterModule.forRoot(appRoutes, { initialNavigation: 'enabledBlocking' })
  ],
  providers: [
    // Services
    LayoutService,
    StyleService,
    PermissionService,
    RequesterService,
    RoleGuardService,

    {
      provide: APP_ENV,
      useValue: environment
    },
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    },
    {
      provide: HTTP_INTERCEPTORS,
      useClass: ErrorsInterceptor,
      multi: true
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
