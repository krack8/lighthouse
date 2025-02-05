import { Component, OnInit, TemplateRef } from '@angular/core';
import { PopoverContent, PopoverRef } from './popover-ref';
import { popoverAnimation } from '@sdk-ui/animations';

@Component({
  templateUrl: './popover.component.html',
  styleUrls: ['./popover.component.scss'],
  animations: [popoverAnimation]
})
export class PopoverComponent implements OnInit {
  renderMethod: 'template' | 'component' | 'text' = 'component';
  content!: PopoverContent;
  context: any;

  constructor(private popoverRef: PopoverRef) {}

  ngOnInit() {
    this.content = this.popoverRef.content;

    if (typeof this.content === 'string') {
      this.renderMethod = 'text';
    }

    if (this.content instanceof TemplateRef) {
      this.renderMethod = 'template';
      this.context = {
        close: this.popoverRef.close.bind(this.popoverRef)
      };
    }
  }
}
