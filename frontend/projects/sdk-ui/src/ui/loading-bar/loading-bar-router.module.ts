import { NgModule } from '@angular/core';
import { RouterModule, Router, NavigationStart, NavigationError, NavigationEnd, NavigationCancel } from '@angular/router';
import { LoadingBarService } from './loading-bar.service';

@NgModule({
  imports: [RouterModule],
  exports: [RouterModule]
})
export class LoadingBarRouterModule {
  constructor(router: Router, loader: LoadingBarService) {
    const ref = loader.useRef('router');
    router.events.subscribe(event => {
      const navState = this.getCurrentNavigationState(router);
      if (navState && navState.ignoreLoadingBar) {
        return;
      }

      if (event instanceof NavigationStart) {
        ref.start();
      }

      if (event instanceof NavigationError || event instanceof NavigationEnd || event instanceof NavigationCancel) {
        ref.complete();
      }
    });
  }

  private getCurrentNavigationState(router: any) {
    // `getCurrentNavigation` only available in angular `7.2`
    const currentNavigation = router.getCurrentNavigation && router.getCurrentNavigation();
    if (currentNavigation && currentNavigation.extras) {
      return currentNavigation.extras.state;
    }

    return {};
  }
}
