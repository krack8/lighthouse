import { ComponentFixture, TestBed } from '@angular/core/testing';
import { LoginComponent } from './login.component';
import { Router } from '@angular/router';
import { ToastrService } from '@sdk-ui/ui';
import { AuthService } from '../auth.service';
import { RequesterService } from '@core-ui/services/requester.service';
import { CoreConfigService } from '@core-ui/services/core-config/core-config.service';
import { UntypedFormBuilder } from '@angular/forms';
import { of, throwError } from 'rxjs';
import { ChangeDetectorRef } from '@angular/core';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;
  let mockRouter = jasmine.createSpyObj('Router', ['navigate']);
  let mockToastr = jasmine.createSpyObj('ToastrService', ['open', 'error', 'warn']);
  let mockAuthService = jasmine.createSpyObj('AuthService', ['mcLogin']);
  let mockRequester = jasmine.createSpyObj('RequesterService', ['get', 'clear']);
  let mockCoreConfigService = { generalInfo$: of({}) };
  let mockCd = jasmine.createSpyObj('ChangeDetectorRef', ['detectChanges']);

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [LoginComponent],
      providers: [
        { provide: Router, useValue: mockRouter },
        { provide: ToastrService, useValue: mockToastr },
        { provide: AuthService, useValue: mockAuthService },
        { provide: RequesterService, useValue: mockRequester },
        { provide: CoreConfigService, useValue: mockCoreConfigService },
        { provide: ChangeDetectorRef, useValue: mockCd },
        UntypedFormBuilder
      ]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should initialize form on ngOnInit', () => {
    component.ngOnInit();
    expect(component.form).toBeDefined();
    expect(component.form.controls['username']).toBeDefined();
    expect(component.form.controls['password']).toBeDefined();
  });

  it('should show error message when form is invalid', () => {
    component.form.setValue({ username: '', password: '' });
    component.submit();
    expect(mockToastr.open).toHaveBeenCalledWith(jasmine.objectContaining({ title: 'Invalid Input' }));
  });

  it('should call authService.mcLogin on valid form submit', () => {
    mockAuthService.mcLogin.and.returnValue(of({ userInfo: { user_is_active: true } }));
    component.form.setValue({ username: 'test@example.com', password: 'password123' });
    component.submit();
    expect(mockAuthService.mcLogin).toHaveBeenCalledWith({ username: 'test@example.com', password: 'password123' });
  });

  it('should handle authentication error correctly', () => {
    mockAuthService.mcLogin.and.returnValue(throwError(() => new Error('Invalid credentials')));
    component.form.setValue({ username: 'test@example.com', password: 'wrongpass' });
    component.submit();
    expect(mockToastr.error).toHaveBeenCalledWith('Invalid credentials');
  });

  it('should navigate to clusters on successful login', () => {
    mockAuthService.mcLogin.and.returnValue(of({ userInfo: { user_is_active: true } }));
    component.form.setValue({ username: 'test@example.com', password: 'password123' });
    component.submit();
    expect(mockRouter.navigate).toHaveBeenCalledWith(['clusters']);
  });

  it('should toggle password visibility', () => {
    expect(component.visible).toBeFalse();
    component.toggleVisibility();
    expect(component.visible).toBeTrue();
    expect(component.inputType).toBe('text');
    component.toggleVisibility();
    expect(component.visible).toBeFalse();
    expect(component.inputType).toBe('password');
  });
});
