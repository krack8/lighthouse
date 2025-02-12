import { inject } from '@angular/core';
import { PageEvent } from '@angular/material/paginator';
import { ActivatedRoute, QueryParamsHandling, Router } from '@angular/router';
import { Subject } from 'rxjs';
import { mapTo, startWith, switchMap, tap } from 'rxjs/operators';

type queryParams = {
  page: number;
  pagelen: number;
  [key: string]: any;
};

type TPaginationConfig = {
  page: number;
  pagelen: number;
  pageSizeOption: number[];
};

export class BasePaginationRoute<T = any> {
  protected route = inject(ActivatedRoute);
  protected router = inject(Router);

  protected _destroyed$ = new Subject<void>();

  protected startingPage!: number;

  protected queryParams!: queryParams;
  protected pageSizeOptions!: number[];
  protected totalElements: number = 0;
  protected elements: T[] = [];

  protected isLoaded = false;
  protected isLoading = true;

  private refreshEvent$ = new Subject<void>();

  protected routeQueryParam$ = this.route.queryParams.pipe(
    switchMap(_qp => {
      const { page, pagelen } = _qp;
      this.queryParams.page = page || this.startingPage;
      if (pagelen) this.queryParams.pagelen = pagelen;
      return this.refreshEvent$.pipe(startWith(undefined), mapTo(_qp));
    })
  );

  constructor(config?: Partial<TPaginationConfig>) {
    this.startingPage = config?.page || 0;
    this.queryParams = {
      page: this.startingPage,
      pagelen: config?.pagelen || 12
    };
    this.pageSizeOptions = config?.pageSizeOption || [12, 24, 36];
  }

  protected updateRouteQueryParams(qp: Partial<queryParams>, queryParamsHandling: QueryParamsHandling = 'merge'): void {
    this.router.navigate([], {
      relativeTo: this.route,
      queryParams: qp,
      queryParamsHandling
    });
  }

  // Pagination
  protected handlePageEvent(event: PageEvent) {
    const qp = {
      page: event.pageIndex,
      pagelen: event.pageSize
    };
    this.updateRouteQueryParams(qp);
  }

  /**
   * @description Refresh with current query param state
   */
  public refresh() {
    this.refreshEvent$.next();
  }
}
