package builders

import (
	"encoding/json"

	gatewayv1beta1 "github.com/kyma-project/api-gateway/api/v1beta1"
	"istio.io/api/security/v1beta1"
	apiv1beta1 "istio.io/api/type/v1beta1"
	securityv1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
)

const (
	istioIngressGatewayPrincipal        string = "cluster.local/ns/istio-system/sa/istio-ingressgateway-service-account"
	oathkeeperMaesterAccountPrincipal   string = "cluster.local/ns/kyma-system/sa/oathkeeper-maester-account"
	authorizationPolicyAppSelectorLabel string = "app"
)

// NewAuthorizationPolicyBuilder returns a builder for istio.io/client-go/pkg/apis/security/v1beta1/AuthorizationPolicy type
func NewAuthorizationPolicyBuilder() *AuthorizationPolicyBuilder {
	return &AuthorizationPolicyBuilder{
		value: &securityv1beta1.AuthorizationPolicy{},
	}
}

type AuthorizationPolicyBuilder struct {
	value *securityv1beta1.AuthorizationPolicy
}

func (ap *AuthorizationPolicyBuilder) Get() *securityv1beta1.AuthorizationPolicy {
	return ap.value
}

func (ap *AuthorizationPolicyBuilder) FromAP(val *securityv1beta1.AuthorizationPolicy) *AuthorizationPolicyBuilder {
	ap.value = val
	return ap
}

func (ap *AuthorizationPolicyBuilder) WithName(val string) *AuthorizationPolicyBuilder {
	ap.value.Name = val
	return ap
}

func (ap *AuthorizationPolicyBuilder) WithGenerateName(val string) *AuthorizationPolicyBuilder {
	ap.value.Name = ""
	ap.value.GenerateName = val
	return ap
}

func (ap *AuthorizationPolicyBuilder) WithNamespace(val string) *AuthorizationPolicyBuilder {
	ap.value.Namespace = val
	return ap
}

func (ap *AuthorizationPolicyBuilder) WithLabel(key, val string) *AuthorizationPolicyBuilder {
	if ap.value.Labels == nil {
		ap.value.Labels = make(map[string]string)
	}
	ap.value.Labels[key] = val
	return ap
}

func (ap *AuthorizationPolicyBuilder) WithSpec(val *v1beta1.AuthorizationPolicy) *AuthorizationPolicyBuilder {
	ap.value.Spec = *val.DeepCopy()
	return ap
}

// NewAuthorizationPolicySpecBuilder returns builder for istio.io/api/security/v1beta1/AuthorizationPolicy type
func NewAuthorizationPolicySpecBuilder() *AuthorizationPolicySpecBuilder {
	return &AuthorizationPolicySpecBuilder{
		value: &v1beta1.AuthorizationPolicy{},
	}
}

type AuthorizationPolicySpecBuilder struct {
	value *v1beta1.AuthorizationPolicy
}

func (aps *AuthorizationPolicySpecBuilder) Get() *v1beta1.AuthorizationPolicy {
	return aps.value
}

func (aps *AuthorizationPolicySpecBuilder) FromAP(val *v1beta1.AuthorizationPolicy) *AuthorizationPolicySpecBuilder {
	aps.value = val
	return aps
}

func (aps *AuthorizationPolicySpecBuilder) WithSelector(val *apiv1beta1.WorkloadSelector) *AuthorizationPolicySpecBuilder {
	aps.value.Selector = val
	return aps
}

func (aps *AuthorizationPolicySpecBuilder) WithRule(val *v1beta1.Rule) *AuthorizationPolicySpecBuilder {
	aps.value.Rules = append(aps.value.Rules, val)
	return aps
}

// NewRuleBuilder returns builder for istio.io/api/security/v1beta1/Rule type
func NewRuleBuilder() *RuleBuilder {
	return &RuleBuilder{
		value: &v1beta1.Rule{},
	}
}

type RuleBuilder struct {
	value *v1beta1.Rule
}

func (r *RuleBuilder) Get() *v1beta1.Rule {
	return r.value
}

func (r *RuleBuilder) WithFrom(val *v1beta1.Rule_From) *RuleBuilder {
	r.value.From = append(r.value.From, val)
	return r
}

func (r *RuleBuilder) WithTo(val *v1beta1.Rule_To) *RuleBuilder {
	r.value.To = append(r.value.To, val)
	return r
}

func (r *RuleBuilder) WithWhenCondition(val *v1beta1.Condition) *RuleBuilder {
	r.value.When = append(r.value.When, val)
	return r
}

// NewFromBuilder returns builder for istio.io/api/security/v1beta1/Rule_From type
func NewFromBuilder() *FromBuilder {
	return &FromBuilder{
		value: &v1beta1.Rule_From{},
	}
}

type FromBuilder struct {
	value *v1beta1.Rule_From
}

func (rf *FromBuilder) Get() *v1beta1.Rule_From {
	return rf.value
}

// WithForcedJWTAuthorization adds RequestPrincipals = "*" requirement, forcing requests to use JWT authorization
func (rf *FromBuilder) WithForcedJWTAuthorization() *FromBuilder {
	// Only support one source at the moment
	source := v1beta1.Source{RequestPrincipals: []string{"*"}}
	rf.value.Source = &source
	return rf
}

func (rf *FromBuilder) WithIngressGatewaySource() *FromBuilder {
	source := v1beta1.Source{Principals: []string{istioIngressGatewayPrincipal}}
	rf.value.Source = &source
	return rf
}

func (rf *FromBuilder) WithOathkeeperProxySource() *FromBuilder {
	source := v1beta1.Source{Principals: []string{oathkeeperMaesterAccountPrincipal}}
	rf.value.Source = &source
	return rf
}

// NewToBuilder returns builder for istio.io/apis/security/v1beta1/Rule_To type
func NewToBuilder() *ToBuilder {
	return &ToBuilder{
		value: &v1beta1.Rule_To{},
	}
}

type ToBuilder struct {
	value *v1beta1.Rule_To
}

func (rt *ToBuilder) Get() *v1beta1.Rule_To {
	return rt.value
}

func (rt *ToBuilder) WithOperation(val *v1beta1.Operation) *ToBuilder {
	rt.value.Operation = val
	return rt
}

// NewOperationBuilder returns builder for istio.io/api/security/v1beta1/Operation type
func NewOperationBuilder() *OperationBuilder {
	return &OperationBuilder{
		value: &v1beta1.Operation{},
	}
}

type OperationBuilder struct {
	value *v1beta1.Operation
}

func (o *OperationBuilder) Get() *v1beta1.Operation {
	return o.value
}

func (o *OperationBuilder) WithMethods(val []string) *OperationBuilder {
	o.value.Methods = val
	return o
}

func (o *OperationBuilder) WithPath(val string) *OperationBuilder {
	o.value.Paths = append(o.value.Paths, val)
	return o
}

// NewConditionBuilder returns builder for istio.io/apis/security/v1beta1/Condition type
func NewConditionBuilder() *ConditionBuilder {
	return &ConditionBuilder{
		value: &v1beta1.Condition{},
	}
}

type ConditionBuilder struct {
	value *v1beta1.Condition
}

func (rc *ConditionBuilder) Get() *v1beta1.Condition {
	return rc.value
}

func (rc *ConditionBuilder) WithKey(key string) *ConditionBuilder {
	rc.value.Key = key
	return rc
}

func (rc *ConditionBuilder) WithValues(values []string) *ConditionBuilder {
	rc.value.Values = values
	return rc
}

// NewRequestAuthenticationBuilder returns a builder for istio.io/client-go/pkg/apis/security/v1beta1/RequestAuthentication type
func NewRequestAuthenticationBuilder() *RequestAuthenticationBuilder {
	return &RequestAuthenticationBuilder{
		value: &securityv1beta1.RequestAuthentication{},
	}
}

type RequestAuthenticationBuilder struct {
	value *securityv1beta1.RequestAuthentication
}

func (ra *RequestAuthenticationBuilder) Get() *securityv1beta1.RequestAuthentication {
	return ra.value
}

func (ra *RequestAuthenticationBuilder) WithFrom(val *securityv1beta1.RequestAuthentication) *RequestAuthenticationBuilder {
	ra.value = val
	return ra
}

func (ra *RequestAuthenticationBuilder) WithName(val string) *RequestAuthenticationBuilder {
	ra.value.Name = val
	return ra
}

func (ra *RequestAuthenticationBuilder) WithGenerateName(val string) *RequestAuthenticationBuilder {
	ra.value.Name = ""
	ra.value.GenerateName = val
	return ra
}

func (ra *RequestAuthenticationBuilder) WithNamespace(val string) *RequestAuthenticationBuilder {
	ra.value.Namespace = val
	return ra
}

func (ra *RequestAuthenticationBuilder) WithLabel(key, val string) *RequestAuthenticationBuilder {
	if ra.value.Labels == nil {
		ra.value.Labels = make(map[string]string)
	}
	ra.value.Labels[key] = val
	return ra
}

func (ra *RequestAuthenticationBuilder) WithSpec(val *v1beta1.RequestAuthentication) *RequestAuthenticationBuilder {
	ra.value.Spec = *val.DeepCopy()
	return ra
}

// NewRequestAuthenticationSpecBuilder returns a builder for istio.io/api/security/v1beta1/RequestAuthentication type
func NewRequestAuthenticationSpecBuilder() *RequestAuthenticationSpecBuilder {
	return &RequestAuthenticationSpecBuilder{
		value: &v1beta1.RequestAuthentication{},
	}
}

type RequestAuthenticationSpecBuilder struct {
	value *v1beta1.RequestAuthentication
}

func (ras *RequestAuthenticationSpecBuilder) Get() *v1beta1.RequestAuthentication {
	return ras.value
}

func (ras *RequestAuthenticationSpecBuilder) WithFrom(val *v1beta1.RequestAuthentication) *RequestAuthenticationSpecBuilder {
	ras.value = val
	return ras
}

func (ras *RequestAuthenticationSpecBuilder) WithSelector(val *apiv1beta1.WorkloadSelector) *RequestAuthenticationSpecBuilder {
	ras.value.Selector = val
	return ras
}

func (ras *RequestAuthenticationSpecBuilder) WithJwtRules(val []*v1beta1.JWTRule) *RequestAuthenticationSpecBuilder {
	ras.value.JwtRules = val
	return ras
}

// NewJwtRuleBuilder returns builder for istio.io/api/security/v1beta1/JWTRule type
func NewJwtRuleBuilder() *JwtRuleBuilder {
	return &JwtRuleBuilder{
		value: &[]*v1beta1.JWTRule{},
	}
}

type JwtRuleBuilder struct {
	value *[]*v1beta1.JWTRule
}

func (jr *JwtRuleBuilder) Get() *[]*v1beta1.JWTRule {
	return jr.value
}

func (jr *JwtRuleBuilder) From(val []*gatewayv1beta1.Authenticator) *JwtRuleBuilder {
	for _, accessStrategy := range val {
		authentications := &Authentications{
			Authentications: []*Authentication{},
		}
		if accessStrategy.Config != nil {
			_ = json.Unmarshal(accessStrategy.Config.Raw, authentications)
		}
		for _, authentication := range authentications.Authentications {
			*jr.value = append(*jr.value, &v1beta1.JWTRule{
				Issuer:  authentication.Issuer,
				JwksUri: authentication.JwksUri,
			})
		}
	}
	return jr
}

// NewSelectorBuilder returns builder for istio.io/api/type/v1beta1/WorkloadSelector type
func NewSelectorBuilder() *SelectorBuilder {
	return &SelectorBuilder{
		value: &apiv1beta1.WorkloadSelector{},
	}
}

type SelectorBuilder struct {
	value *apiv1beta1.WorkloadSelector
}

func (s *SelectorBuilder) Get() *apiv1beta1.WorkloadSelector {
	return s.value
}

func (s *SelectorBuilder) WithMatchLabels(key, val string) *SelectorBuilder {
	if s.value.MatchLabels == nil {
		s.value.MatchLabels = make(map[string]string)
	}
	s.value.MatchLabels[key] = val
	return s
}

type Authentications struct {
	Authentications []*Authentication `json:"authentications"`
}

type Authentication struct {
	Issuer  string `json:"issuer"`
	JwksUri string `json:"jwksUri"`
}

func SelectorFromService(service *gatewayv1beta1.Service) *apiv1beta1.WorkloadSelector {
	return &apiv1beta1.WorkloadSelector{
		MatchLabels: map[string]string{authorizationPolicyAppSelectorLabel: *service.Name},
	}
}
