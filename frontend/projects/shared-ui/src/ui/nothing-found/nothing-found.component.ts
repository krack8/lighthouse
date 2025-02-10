import { NgIf } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  standalone: true,
  selector: 'kc-nothing-found',
  imports: [NgIf],
  templateUrl: './nothing-found.component.html',
  styleUrls: ['./nothing-found.component.scss']
})
export class NothingFoundComponent {
  @Input() title!: string;
  @Input() message!: string;
}
