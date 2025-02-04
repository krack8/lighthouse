import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit, OnDestroy } from '@angular/core';
import { UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import icVisibility from '@iconify/icons-ic/twotone-visibility';
import icVisibilityOff from '@iconify/icons-ic/twotone-visibility-off';
import { fadeInUp400ms } from '@sdk-ui/animations/fade-in-up.animation';
import { ToastrService } from '@sdk-ui/ui';
import { AuthService } from '../auth.service';
import { RequesterService } from '@core-ui/services/requester.service';
import { CoreConfigService } from '@core-ui/services/core-config/core-config.service';
import { Subject } from 'rxjs';
import { ICoreConfig } from '@core-ui/services/core-config/core-config.interfaces';

@Component({
  selector: 'kc-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
  animations: [fadeInUp400ms]
})
export class LoginComponent implements OnInit, OnDestroy {
  private _destroy$: Subject<void> = new Subject<void>();
  form: UntypedFormGroup;

  inputType = 'password';
  visible = false;

  icVisibility = icVisibility;
  icVisibilityOff = icVisibilityOff;

  errorMessage: string;
  isSubmitting!: boolean;

  currentUser: any;
  auth2: any;

  coreConfig!: ICoreConfig;

  isSsoInfoLoading: boolean = false;

  constructor(
    private router: Router,
    private toastr: ToastrService,
    private fb: UntypedFormBuilder,
    private cd: ChangeDetectorRef,
    private authService: AuthService,
    private requester: RequesterService,
    private coreConfigService: CoreConfigService
  ) {}

  ngOnInit() {
    this.coreConfigService.generalInfo$.subscribe((config: ICoreConfig) => {
      this.coreConfig = config;
    });
    this.currentUser = this.requester.get();
    if (this.currentUser != null) {
      this.goToNextPage(this.currentUser?.userInfo);
    }

    this.form = this.fb.group({
      username: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]]
    });
  }

  ngOnDestroy(): void {
    this._destroy$.next();
    this._destroy$.complete();
  }

  send() {
    this.router.navigate(['/']);
  }

  submit() {
    if (this.form.invalid) {
      this.toastr.open({
        icon: 'icon-error',
        iconBg: 'text-error',
        title: 'Invalid Input',
        message: 'Please fill the required fields'
      });
      return;
    }

    if (!this.isSubmitting) {
      const payload = { username: this.form.value.username, password: this.form.value.password };

      this.isSubmitting = true;
      this.cd.detectChanges();

      this.authService.mcLogin(payload).subscribe({
        next: (profileData: any) => {
          this.goToNextPage(profileData);
          this.isSubmitting = false;
        },
        error: err => {
          console.log(err);
          this.toastr.error(err.message);
          this.cd.detectChanges();
          this.isSubmitting = false;
        },
        complete: () => {
          this.cd.detectChanges();
        }
      });
    } else {
      this.toastr.open({
        icon: 'icon-warn',
        iconBg: 'text-warn',
        title: 'Please Wait',
        message: 'A sign in request is already in progress!'
      });
    }
  }

  goToNextPage(userData: any) {
    // if (userData?.userInfo?.user_is_active === false) {
    //   if (userData?.userInfo?.user_type === 'NON_ADMIN') {
    //     this.toastr.warn('Your account is deactivated. Please contact your admin', 'Activation Required');
    //   } else {
    //     this.toastr.warn('Your account is deactivated', 'Activation Required');
    //   }
    //   this.requester.clear();
    //   this.router.navigate(['/auth/login']);
    //   return;
    // }

    this.router.navigate(['clusters']);
    return;
  }

  toggleVisibility() {
    if (this.visible) {
      this.inputType = 'password';
      this.visible = false;
      this.cd.detectChanges();
    } else {
      this.inputType = 'text';
      this.visible = true;
      this.cd.detectChanges();
    }
  }
}
