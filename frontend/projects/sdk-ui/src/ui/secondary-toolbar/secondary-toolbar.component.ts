import { Component, Input } from '@angular/core';

@Component({
  selector: 'kc-secondary-toolbar',
  templateUrl: './secondary-toolbar.component.html',
  styleUrls: ['./secondary-toolbar.component.scss']
})
export class SecondaryToolbarComponent {
  @Input() current!: string;
  @Input() crumbs!: string[];
}
