import { ActivatedRouteSnapshot } from '@angular/router';
import { KcRouteData } from '@sdk-ui/interfaces';

export function checkRouterChildsData(
  route: ActivatedRouteSnapshot & { data?: KcRouteData },
  compareWith: (data: KcRouteData) => boolean
): boolean {
  if (compareWith(route.data)) return true;

  if (!route.firstChild) return false;

  return checkRouterChildsData(route.firstChild, compareWith);
}
