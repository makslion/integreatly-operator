package common

/*func rhmi2Namespaces() []string {
	return []string{
		MonitoringOperatorNamespace,
		MonitoringFederateNamespace,
		AMQOnlineOperatorNamespace,
		ApicuritoProductNamespace,
		ApicuritoOperatorNamespace,
		CloudResourceOperatorNamespace,
		CodeReadyProductNamespace,
		CodeReadyOperatorNamespace,
		FuseProductNamespace,
		FuseOperatorNamespace,
		RHSSOUserOperatorNamespace,
		RHSSOProductNamespace,
		RHSSOOperatorNamespace,
		SolutionExplorerProductNamespace,
		SolutionExplorerOperatorNamespace,
		ThreeScaleProductNamespace,
		ThreeScaleOperatorNamespace,
		UPSProductNamespace,
		UPSOperatorNamespace,
	}
}

func managedApiNamespaces() []string {
	return []string{
		MonitoringOperatorNamespace,
		CloudResourceOperatorNamespace,
		RHSSOUserOperatorNamespace,
		RHSSOProductNamespace,
		RHSSOOperatorNamespace,
		ThreeScaleProductNamespace,
		ThreeScaleOperatorNamespace,
		Marin3rOperatorNamespace,
		Marin3rProductNamespace,
		CustomerGrafanaNamespace,
	}
}*/

func TestNamespaceCreated(t TestingTB, ctx *TestingContext) {

	/*	namespacesCreated := getNamespaces(t, ctx)

		for _, namespace := range namespacesCreated {
			ns := &corev1.Namespace{}
			err := ctx.Client.Get(goctx.TODO(), k8sclient.ObjectKey{Name: namespace}, ns)

			if err != nil {
				t.Errorf("Expected %s namespace to be created but wasn't: %s", namespace, err)
				continue
			}
		}*/
}

/*func getNamespaces(t TestingTB, ctx *TestingContext) []string {

	//get RHMI
	rhmi, err := GetRHMI(ctx.Client, true)
	if err != nil {
		t.Errorf("error getting RHMI CR: %v", err)
	}

	if rhmi.Spec.Type == string(integreatlyv1alpha1.InstallationTypeManagedApi) {
		return managedApiNamespaces()
	} else {
		return rhmi2Namespaces()
	}
}*/
