import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { SelectedClusterService } from '@core-ui/services/selected-cluster.service';
import { NavigationDropdown, NavigationItem, NavigationLink, NavigationSubheading } from '@sdk-ui/interfaces';
import { BehaviorSubject, Observable, Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class NavigationService {
  private _destroy$ = new Subject<void>();
  private navigationItems = new BehaviorSubject<NavigationItem[]>([]);
  Items$: Observable<NavigationItem[]> = this.navigationItems.asObservable();

  private _openChangeSubject = new Subject<NavigationDropdown>();
  openChange$ = this._openChangeSubject.asObservable();

  constructor(private selectedClusterService: SelectedClusterService, private router: Router,) {}

  loadItems(items: NavigationItem[]): void {
    this.navigationItems.next(this.initializeClusterId(items));
  }

  get items(): NavigationItem[] {
    return this.navigationItems.value;
  }

  // Child

  triggerOpenChange(item: NavigationDropdown) {
    this._openChangeSubject.next(item);
  }

  isLink(item: NavigationItem): item is NavigationLink {
    return item.type === 'link';
  }

  isDropdown(item: NavigationItem): item is NavigationDropdown {
    return item.type === 'dropdown';
  }

  isSubheading(item: NavigationItem): item is NavigationSubheading {
    return item.type === 'subheading';
  }

  initializeClusterId(items: NavigationItem[]): NavigationItem[] {
    this.selectedClusterService.clusterId$.pipe(
     takeUntil(this._destroy$)
    ).subscribe(clusterId => {
      const replaceClusterId = (navItem: NavigationItem, search: string, replace: string) => {
        if (navItem.route?.includes(search)) {
          navItem.route = navItem.route.replace(search, replace);
        }
        if (navItem.children) {
          navItem.children.forEach(child => replaceClusterId(child, search, replace));
        }
      };
  
      items.forEach(item => {
        if (clusterId) {
          replaceClusterId(item, ':clusterId', clusterId);
        }
      });
    });

    return items;
  }

}
