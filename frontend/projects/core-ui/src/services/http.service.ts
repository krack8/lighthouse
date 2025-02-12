import { Inject, Injectable, Optional } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { APP_ENV } from '@core-ui/constants';
import { IAppEnv } from '@core-ui/interfaces';

/**
 * @description Http Service is MultiCluster http entrypoint
 */
@Injectable({
  providedIn: 'root'
})
export class HttpService {
  apiBaseUrl: string;

  constructor(
    private http: HttpClient,
    @Optional() @Inject(APP_ENV) private _env: IAppEnv
  ) {
    this.apiBaseUrl = _env?.apiEndPoint;
  }

  get(url: string, queryParams?: any, responseType?: any): Observable<any> {
    const queryParameters = queryParams ? queryParams : {};
    const responseTypes = responseType ? responseType : null;

    const httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: queryParameters,
      responseType: responseTypes
    };

    return this.http.get(this.apiBaseUrl + url, httpOptions);
  }

  post(url: string, data: any, queryParams?: any): Observable<any> {
    const httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: queryParams || {}
    };
    const body = JSON.stringify(data);
    return this.http.post(this.apiBaseUrl + url, body, httpOptions);
  }

  put(url: string, data: any, queryParams?: any): Observable<any> {
    const queryParameters = queryParams ? queryParams : {};
    const httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: queryParameters
    };
    const body = JSON.stringify(data);
    return this.http.put(this.apiBaseUrl + url, body, httpOptions);
  }

  delete(url: string, queryParams?: any): Observable<any> {
    const queryParameters = queryParams ? queryParams : {};
    const httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: queryParameters
    };
    return this.http.delete(this.apiBaseUrl + url, httpOptions);
  }

  upload(url: string, payload: any): Observable<any> {
    const httpOptions = {
      headers: new HttpHeaders({ Accept: 'application/json' })
    };
    const body = JSON.stringify(payload);
    return this.http.post(this.apiBaseUrl + url, body, httpOptions);
  }
}
