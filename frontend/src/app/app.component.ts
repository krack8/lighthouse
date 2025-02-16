import { Platform } from '@angular/cdk/platform';
import { DOCUMENT } from '@angular/common';
import { Component, Inject, OnInit, Renderer2 } from '@angular/core';
import { MatIconRegistry } from '@angular/material/icon';
import { Title } from '@angular/platform-browser';
import { ActivatedRoute, Router, NavigationEnd } from '@angular/router';
import { filter, map, mergeMap, take } from 'rxjs/operators';
import { StyleService, Style, SplashScreenService } from '@sdk-ui/services';

@Component({
  selector: 'kc-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'kc';

  constructor(
    @Inject(DOCUMENT) private document: Document,
    private styleService: StyleService,
    private iconRegistry: MatIconRegistry,
    private renderer: Renderer2,
    private platform: Platform,
    private route: ActivatedRoute,
    private router: Router,
    private titleService: Title,
    private splashScreenService: SplashScreenService
  ) {
    // Remove initilize loader
    this.router.events
      .pipe(
        filter(event => event instanceof NavigationEnd),
        take(1)
      )
      .subscribe(() => {
        if (this.splashScreenService.splashScreenElem) {
          this.splashScreenService.hide();
        }
      });
    this.styleService.setStyle(Style.dark);

    this.iconRegistry.setDefaultFontSetClass('iconify');

    if (this.platform.BLINK) {
      this.renderer.addClass(this.document.body, 'is-blink');
    }
  }

  ngOnInit(): void {
    this.router.events
      .pipe(
        filter(event => event instanceof NavigationEnd),
        map(() => this.route),
        map(route => {
          while (route.firstChild) {
            route = route.firstChild;
          }
          return route;
        }),
        filter(route => route.outlet === 'primary'),
        mergeMap(route => route.data)
      )
      .subscribe(event => {
        this.titleService.setTitle('Lighthouse' + ' | ' + event['title']);
      });
  }
}
