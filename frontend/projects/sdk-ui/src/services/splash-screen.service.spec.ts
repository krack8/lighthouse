import { TestBed } from '@angular/core/testing';

import { SplashScreenService } from './splash-screen.service';

describe('SplashScreenService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: SplashScreenService = TestBed.get(SplashScreenService);
    expect(service).toBeTruthy();
  });
});
