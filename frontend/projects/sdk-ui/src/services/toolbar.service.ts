import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

type TToolbar = {
  title: string;
};

@Injectable({
  providedIn: 'root'
})
export class ToolbarService {
  private data = new BehaviorSubject<TToolbar>({ title: '' });
  currentData = this.data.asObservable();

  changeData(data: TToolbar) {
    this.data.next(data);
  }
}
