import { Component, ElementRef, EventEmitter, Input, OnChanges, Optional, Output, SimpleChanges, ViewChild } from '@angular/core';
import { CoreConfigService } from '@core-ui/services/core-config/core-config.service';
import { Ace, edit } from 'ace-builds';

import 'ace-builds';
import 'ace-builds/src-noconflict/theme-tomorrow_night_blue';

@Component({
  selector: 'kc-ace-editor',
  templateUrl: './ace-editor.component.html',
  styleUrls: ['./ace-editor.component.scss'],
  host: {
    class: 'kc-ace-editor-host'
  }
})
export class AceEditorComponent implements OnChanges {
  @ViewChild('editor') editorRef!: ElementRef;
  @Output() textChange = new EventEmitter<string>();
  @Input() text!: string;
  @Input() readOnly: boolean = false;
  @Input() mode: string = 'json';
  @Input() prettify: boolean = true;

  private theme: string;

  editor!: Ace.Editor;

  // All possible options can be found at:
  // https://github.com/ajaxorg/ace/wiki/Configuring-Ace
  options = {
    //? renderer
    highlightActiveLine: true,
    vScrollBarAlwaysVisible: true,
    displayIndentGuides: true,
    fontSize: '14px',

    //? session
    tabSize: 2,
    wrap: true
  };

  constructor(@Optional() public coreConfigService: CoreConfigService) {}

  ngOnChanges(changes: SimpleChanges): void {
    if (!this.editor) {
      return;
    }

    for (const propName in changes) {
      if (changes.hasOwnProperty(propName)) {
        switch (propName) {
          case 'text':
            this.onExternalUpdate_();
            break;
          case 'mode':
            this.onEditorModeChange_();
            break;
          default:
        }
      }
    }
  }

  ngAfterViewInit(): void {
    this.theme = this.coreConfigService?.generalInfoSnapshot?.webTheme || 'DARK';
    // setTimeout solve editor cursor problem
    setTimeout(() => {
      this.initEditor_();
    }, 100);
  }

  /**
   * @Definition Emit editor value
   * @Params text: string
   */
  onTextChange(text: string): void {
    this.textChange.emit(text);
  }

  /**
   * @Definition Initialize Editor
   */
  private initEditor_(): void {
    this.editor = edit(this.editorRef.nativeElement);
    this.editor.setOptions(this.options);
    this.editor.setValue(this.text, -1);
    this.editor.setReadOnly(this.readOnly);

    if (this.theme === 'DARK') {
      this.editor.setTheme('ace/theme/tomorrow_night_blue');
    } else {
      this.editor.setTheme('ace/theme/chrome');
    }

    this.setEditorMode_();
    this.editor.session.setUseWorker(false);
    this.editor.on('change', () => this.onEditorTextChange_());
  }

  /**
   * @Definition Update when User input new key
   */
  private onExternalUpdate_(): void {
    const point = this.editor.getCursorPosition();
    this.editor.setValue(this.text, -1);
    this.editor.moveCursorToPosition(point);
  }

  /**
   * @Definition callback for data change
   */
  private onEditorTextChange_(): void {
    this.text = this.editor.getValue();
    this.onTextChange(this.text);
  }

  /**
   * @Definition Update editor mode when change
   */
  private onEditorModeChange_(): void {
    this.setEditorMode_();
  }

  /**
   * @Definition Update editor mode
   */
  private setEditorMode_(): void {
    this.editor.getSession().setMode(`ace/mode/${this.mode}`);
  }
}
