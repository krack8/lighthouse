import { Component, OnInit } from '@angular/core';
import { ICoreConfig } from '@core-ui/services/core-config/core-config.interfaces';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { CoreConfigService } from '@core-ui/services/core-config/core-config.service';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';

@Component({
  selector: 'kc-system-settings',
  templateUrl: './system-settings.component.html',
  styleUrls: ['./system-settings.component.scss']
})
export class SystemSettingsComponent implements OnInit {
  private _destroy$ = new Subject<void>();

  coreConfig!: ICoreConfig;

  constructor(
    private coreConfigService: CoreConfigService,
    private _toolbarService: ToolbarService
  ) {
    this._toolbarService.changeData({ title: 'Settings' });
  }

  ngOnInit(): void {
    this.coreConfigService.generalInfo$.pipe(takeUntil(this._destroy$)).subscribe((config: ICoreConfig) => {
      this.coreConfig = config;
    });
  }
}
