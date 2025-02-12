# API reference for Tooltip (HTML Structure tooltip)

```ts
import { CdkTooltipDirective } from '@cdk-ui/tooltip';
```

## Usage

#### component.html

```html
<div [cdkTooltip]="tooltipTemplate">Hover me for tooltip</div>

<ng-template #tooltipTemplate>
  <strong>Hello K.</strong>
  <p>Nice to meet you</p>
</ng-template>
```
