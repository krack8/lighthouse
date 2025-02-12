import { HttpClient, HttpClientModule } from '@angular/common/http';
import { APP_INITIALIZER, NgModule } from '@angular/core';
import { CoreConfigService } from '@core-ui/services';

@NgModule({
  imports: [HttpClientModule],
  providers: [
    {
      provide: APP_INITIALIZER,
      useFactory: (coreConfigService: CoreConfigService) => {
        return () => coreConfigService.loadConfigurationData();
      },
      deps: [CoreConfigService, HttpClient],
      multi: true
    }
  ]
})
export class InitializerModule {}
