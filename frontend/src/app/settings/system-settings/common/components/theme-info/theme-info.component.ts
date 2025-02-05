import { Component, Input, OnInit } from '@angular/core';
import { fadeInUp400ms } from '@sdk-ui/animations/fade-in-up.animation';
import { fadeInRight400ms } from '@sdk-ui/animations/fade-in-right.animation';
import { ICoreConfig, Theme } from '@core-ui/services/core-config/core-config.interfaces';
import { CoreConfigService } from '@core-ui/services/core-config/core-config.service';
import { ToastrService } from '@sdk-ui/ui';
import { RequesterService } from '@core-ui/services/requester.service';
import { SystemSettingsService } from '@settings/system-settings/system-settings.service';

@Component({
  selector: 'kc-theme-info',
  templateUrl: './theme-info.component.html',
  styleUrls: ['./theme-info.component.scss'],
  animations: [fadeInUp400ms, fadeInRight400ms]
})
export class ThemeInfoComponent implements OnInit {
  isThemeInfo = false;
  selectedTheme!: Theme;
  isSubmitting: boolean = false;

  user: any;

  @Input() set config(value: ICoreConfig) {
    this.selectedTheme = value?.webTheme || Theme.DARK;
  }

  constructor(
    private coreConfService: CoreConfigService,
    private toastr: ToastrService,
    private _requesterService: RequesterService,
    private _settingsService: SystemSettingsService
  ) {}

  ngOnInit(): void {
    this.user = this._requesterService.get();
  }

  showThemeInfo() {
    this.isThemeInfo = !this.isThemeInfo;
  }

  selectTheme(theme) {
    this.selectedTheme = theme;
  }

  saveThemeInfo() {
    this.isSubmitting = true;
    const qp = {
      webTheme: this.selectedTheme
    };
    this._settingsService.saveWebThemePreference(qp).subscribe({
      next: res => {
        if (res.status === 'success') {
          this.isSubmitting = false;

          this.user.userInfo.webThemePreference = this.selectedTheme;

          if (this.selectedTheme === 'DARK') {
            this.coreConfService.updateTheme(Theme.DARK);
          } else if (this.selectedTheme === 'LIGHT_PINK') {
            this.coreConfService.updateTheme(Theme.LIGHT_PINK);
          } else {
            this.coreConfService.updateTheme(Theme.LIGHT);
          }

          this._requesterService.save(this.user);

          this.toastr.success('Theme preference saved');
        } else {
          this.isSubmitting = false;
          this.toastr.error('Failed saving theme preference. Please try again!');
        }
      },
      error: () => {
        this.isSubmitting = false;
        this.toastr.error('Failed saving theme preference. Please try again!');
      }
    });
  }
}
