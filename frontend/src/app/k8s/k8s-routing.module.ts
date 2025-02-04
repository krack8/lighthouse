import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { K8sComponent } from './k8s.component';
import { K8sResolver } from './k8s.resolver';

const k8sRoutes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'node-list'
  },
  {
    path: 'namespaces',
    loadChildren: () => import('./k8s-namespaces/k8s-namespaces.module').then(m => m.K8sNamespacesModule),
    data: {
      permissions: ['VIEW_K8S_NAMESPACE']
    }
  },
  {
    path: 'node-list',
    loadChildren: () => import('./k8s-nodes/k8s-nodes.module').then(m => m.K8sNodesModule),
    data: {
      permissions: ['VIEW_K8S_NODES']
    }
  },
  {
    path: 'persistent-volume',
    loadChildren: () => import('./k8s-persistent-volume/k8s-persistent-volume.module').then(m => m.K8sPersistentVolumeModule),
    data: {
      permissions: ['VIEW_K8S_PERSISTENT_VOLUME']
    }
  },
  {
    path: 'cluster-role',
    loadChildren: () => import('./k8s-cluster-role/k8s-cluster-role.module').then(m => m.K8sClusterRoleModule),
    data: {
      permissions: ['VIEW_K8S_CLUSTER_ROLE']
    }
  },
  {
    path: 'cluster-role-binding',
    loadChildren: () => import('./k8s-cluster-role-binding/k8s-cluster-role-binding.module').then(m => m.K8sClusterRoleBindingModule),
    data: {
      permissions: ['VIEW_K8S_CLUSTER_ROLE_BINDING']
    }
  },
  {
    path: 'storage-class',
    loadChildren: () => import('./k8s-storage-class/k8s-storage-class.module').then(m => m.K8sStorageClassModule),
    data: {
      permissions: ['VIEW_K8S_STORAGE_CLASS']
    }
  },
  {
    path: 'custom-resources-defination',
    loadChildren: () =>
      import('./k8s-cluster-custom-resources/k8s-cluster-custom-resources.module').then(m => m.K8sClusterCustomResourcesModule),
    data: {
      permissions: ['VIEW_K8S_CUSTOM_RESOURCE_DEFINATION']
    }
  },
  {
    path: 'terminal',
    loadChildren: () => import('./k8s-terminal/k8s-terminal.module').then(m => m.K8sTerminalModule)
  }
];

const routes: Routes = [
  {
    path: '',
    component: K8sComponent,
    resolve: { cluster: K8sResolver },
    children: k8sRoutes
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class K8sRoutingModule {}
