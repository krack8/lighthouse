import { ChangeDetectionStrategy, Component, Input } from '@angular/core';
import { MatRipple } from '@angular/material/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Clipboard } from '@angular/cdk/clipboard';

@Component({
  selector: 'cdk-clipboard',
  changeDetection: ChangeDetectionStrategy.OnPush,
  providers: [MatRipple],
  template: `<svg class="cdk-clipboard__icon" width="28" height="28" viewBox="0 0 28 28" fill="none" xmlns="http://www.w3.org/2000/svg">
    <circle class="cdk-clipboard__icon__bg" cx="14" cy="14" r="14" />
    <path
      class="cdk-clipboard__icon__color"
      d="M20 11V20H11V11H20ZM20 10H11C10.7348 10 10.4804 10.1054 10.2929 10.2929C10.1054 10.4804 10 10.7348 10 11V20C10 20.2652 10.1054 20.5196 10.2929 20.7071C10.4804 20.8946 10.7348 21 11 21H20C20.2652 21 20.5196 20.8946 20.7071 20.7071C20.8946 20.5196 21 20.2652 21 20V11C21 10.7348 20.8946 10.4804 20.7071 10.2929C20.5196 10.1054 20.2652 10 20 10Z"
    />
    <path
      class="cdk-clipboard__icon__color"
      d="M8 15H7V8C7 7.73478 7.10536 7.48043 7.29289 7.29289C7.48043 7.10536 7.73478 7 8 7H15V8H8V15Z"
    />
  </svg> `,
  host: {
    class: 'cdk-clipboard',
    '(click)': 'copy($event)'
  }
})
export class CdkClipboardComponent {
  @Input() cbContent!: string;
  @Input() message?: string;

  constructor(
    private _clipboard: Clipboard,
    private snackbar: MatSnackBar,
    private ripple: MatRipple
  ) {}

  copy(e: PointerEvent): void {
    this._clipboard.copy(this.cbContent);
    this.snackbar.open(this.message || 'Copied to the clipboard!', 'Close', {
      duration: 3000,
      panelClass: ['snackbar-dark']
    });
    this.ripple.launch(e.x, e.y);
  }
}
