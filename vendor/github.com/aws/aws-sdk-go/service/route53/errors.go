// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

const (

	// ErrCodeConcurrentModification for service response error code
	// "ConcurrentModification".
	//
	// Another user submitted a request to update the object at the same time that
	// you did. Retry the request.
	ErrCodeConcurrentModification = "ConcurrentModification"

	// ErrCodeConflictingDomainExists for service response error code
	// "ConflictingDomainExists".
	//
	// You specified an Amazon VPC that you're already using for another hosted
	// zone, and the domain that you specified for one of the hosted zones is a
	// subdomain of the domain that you specified for the other hosted zone. For
	// example, you can't use the same Amazon VPC for the hosted zones for example.com
	// and test.example.com.
	ErrCodeConflictingDomainExists = "ConflictingDomainExists"

	// ErrCodeConflictingTypes for service response error code
	// "ConflictingTypes".
	//
	// You tried to update a traffic policy instance by using a traffic policy version
	// that has a different DNS type than the current type for the instance. You
	// specified the type in the JSON document in the CreateTrafficPolicy or CreateTrafficPolicyVersionrequest.
	ErrCodeConflictingTypes = "ConflictingTypes"

	// ErrCodeDelegationSetAlreadyCreated for service response error code
	// "DelegationSetAlreadyCreated".
	//
	// A delegation set with the same owner and caller reference combination has
	// already been created.
	ErrCodeDelegationSetAlreadyCreated = "DelegationSetAlreadyCreated"

	// ErrCodeDelegationSetAlreadyReusable for service response error code
	// "DelegationSetAlreadyReusable".
	//
	// The specified delegation set has already been marked as reusable.
	ErrCodeDelegationSetAlreadyReusable = "DelegationSetAlreadyReusable"

	// ErrCodeDelegationSetInUse for service response error code
	// "DelegationSetInUse".
	//
	// The specified delegation contains associated hosted zones which must be deleted
	// before the reusable delegation set can be deleted.
	ErrCodeDelegationSetInUse = "DelegationSetInUse"

	// ErrCodeDelegationSetNotAvailable for service response error code
	// "DelegationSetNotAvailable".
	//
	// You can create a hosted zone that has the same name as an existing hosted
	// zone (example.com is common), but there is a limit to the number of hosted
	// zones that have the same name. If you get this error, Amazon Route 53 has
	// reached that limit. If you own the domain name and Amazon Route 53 generates
	// this error, contact Customer Support.
	ErrCodeDelegationSetNotAvailable = "DelegationSetNotAvailable"

	// ErrCodeDelegationSetNotReusable for service response error code
	// "DelegationSetNotReusable".
	//
	// A reusable delegation set with the specified ID does not exist.
	ErrCodeDelegationSetNotReusable = "DelegationSetNotReusable"

	// ErrCodeHealthCheckAlreadyExists for service response error code
	// "HealthCheckAlreadyExists".
	//
	// The health check you're attempting to create already exists. Amazon Route
	// 53 returns this error when a health check has already been created with the
	// specified value for CallerReference.
	ErrCodeHealthCheckAlreadyExists = "HealthCheckAlreadyExists"

	// ErrCodeHealthCheckInUse for service response error code
	// "HealthCheckInUse".
	//
	// The health check ID for this health check is referenced in the HealthCheckId
	// element in one of the resource record sets in one of the hosted zones that
	// are owned by the current AWS account.
	ErrCodeHealthCheckInUse = "HealthCheckInUse"

	// ErrCodeHealthCheckVersionMismatch for service response error code
	// "HealthCheckVersionMismatch".
	//
	// The value of HealthCheckVersion in the request doesn't match the value of
	// HealthCheckVersion in the health check.
	ErrCodeHealthCheckVersionMismatch = "HealthCheckVersionMismatch"

	// ErrCodeHostedZoneAlreadyExists for service response error code
	// "HostedZoneAlreadyExists".
	//
	// The hosted zone you're trying to create already exists. Amazon Route 53 returns
	// this error when a hosted zone has already been created with the specified
	// CallerReference.
	ErrCodeHostedZoneAlreadyExists = "HostedZoneAlreadyExists"

	// ErrCodeHostedZoneNotEmpty for service response error code
	// "HostedZoneNotEmpty".
	//
	// The hosted zone contains resource records that are not SOA or NS records.
	ErrCodeHostedZoneNotEmpty = "HostedZoneNotEmpty"

	// ErrCodeHostedZoneNotFound for service response error code
	// "HostedZoneNotFound".
	//
	// The specified HostedZone can't be found.
	ErrCodeHostedZoneNotFound = "HostedZoneNotFound"

	// ErrCodeIncompatibleVersion for service response error code
	// "IncompatibleVersion".
	//
	// The resource you're trying to access is unsupported on this Amazon Route
	// 53 endpoint.
	ErrCodeIncompatibleVersion = "IncompatibleVersion"

	// ErrCodeInvalidArgument for service response error code
	// "InvalidArgument".
	//
	// Parameter name is invalid.
	ErrCodeInvalidArgument = "InvalidArgument"

	// ErrCodeInvalidChangeBatch for service response error code
	// "InvalidChangeBatch".
	//
	// This exception contains a list of messages that might contain one or more
	// error messages. Each error message indicates one error in the change batch.
	ErrCodeInvalidChangeBatch = "InvalidChangeBatch"

	// ErrCodeInvalidDomainName for service response error code
	// "InvalidDomainName".
	//
	// The specified domain name is not valid.
	ErrCodeInvalidDomainName = "InvalidDomainName"

	// ErrCodeInvalidInput for service response error code
	// "InvalidInput".
	//
	// The input is not valid.
	ErrCodeInvalidInput = "InvalidInput"

	// ErrCodeInvalidPaginationToken for service response error code
	// "InvalidPaginationToken".
	ErrCodeInvalidPaginationToken = "InvalidPaginationToken"

	// ErrCodeInvalidTrafficPolicyDocument for service response error code
	// "InvalidTrafficPolicyDocument".
	//
	// The format of the traffic policy document that you specified in the Document
	// element is invalid.
	ErrCodeInvalidTrafficPolicyDocument = "InvalidTrafficPolicyDocument"

	// ErrCodeInvalidVPCId for service response error code
	// "InvalidVPCId".
	//
	// The VPC ID that you specified either isn't a valid ID or the current account
	// is not authorized to access this VPC.
	ErrCodeInvalidVPCId = "InvalidVPCId"

	// ErrCodeLastVPCAssociation for service response error code
	// "LastVPCAssociation".
	//
	// The VPC that you're trying to disassociate from the private hosted zone is
	// the last VPC that is associated with the hosted zone. Amazon Route 53 doesn't
	// support disassociating the last VPC from a hosted zone.
	ErrCodeLastVPCAssociation = "LastVPCAssociation"

	// ErrCodeLimitsExceeded for service response error code
	// "LimitsExceeded".
	//
	// The limits specified for a resource have been exceeded.
	ErrCodeLimitsExceeded = "LimitsExceeded"

	// ErrCodeNoSuchChange for service response error code
	// "NoSuchChange".
	//
	// A change with the specified change ID does not exist.
	ErrCodeNoSuchChange = "NoSuchChange"

	// ErrCodeNoSuchDelegationSet for service response error code
	// "NoSuchDelegationSet".
	//
	// A reusable delegation set with the specified ID does not exist.
	ErrCodeNoSuchDelegationSet = "NoSuchDelegationSet"

	// ErrCodeNoSuchGeoLocation for service response error code
	// "NoSuchGeoLocation".
	//
	// Amazon Route 53 doesn't support the specified geolocation.
	ErrCodeNoSuchGeoLocation = "NoSuchGeoLocation"

	// ErrCodeNoSuchHealthCheck for service response error code
	// "NoSuchHealthCheck".
	//
	// No health check exists with the ID that you specified in the DeleteHealthCheck
	// request.
	ErrCodeNoSuchHealthCheck = "NoSuchHealthCheck"

	// ErrCodeNoSuchHostedZone for service response error code
	// "NoSuchHostedZone".
	//
	// No hosted zone exists with the ID that you specified.
	ErrCodeNoSuchHostedZone = "NoSuchHostedZone"

	// ErrCodeNoSuchTrafficPolicy for service response error code
	// "NoSuchTrafficPolicy".
	//
	// No traffic policy exists with the specified ID.
	ErrCodeNoSuchTrafficPolicy = "NoSuchTrafficPolicy"

	// ErrCodeNoSuchTrafficPolicyInstance for service response error code
	// "NoSuchTrafficPolicyInstance".
	//
	// No traffic policy instance exists with the specified ID.
	ErrCodeNoSuchTrafficPolicyInstance = "NoSuchTrafficPolicyInstance"

	// ErrCodeNotAuthorizedException for service response error code
	// "NotAuthorizedException".
	//
	// Associating the specified VPC with the specified hosted zone has not been
	// authorized.
	ErrCodeNotAuthorizedException = "NotAuthorizedException"

	// ErrCodePriorRequestNotComplete for service response error code
	// "PriorRequestNotComplete".
	//
	// If Amazon Route 53 can't process a request before the next request arrives,
	// it will reject subsequent requests for the same hosted zone and return an
	// HTTP 400 error (Bad request). If Amazon Route 53 returns this error repeatedly
	// for the same request, we recommend that you wait, in intervals of increasing
	// duration, before you try the request again.
	ErrCodePriorRequestNotComplete = "PriorRequestNotComplete"

	// ErrCodePublicZoneVPCAssociation for service response error code
	// "PublicZoneVPCAssociation".
	//
	// You're trying to associate a VPC with a public hosted zone. Amazon Route
	// 53 doesn't support associating a VPC with a public hosted zone.
	ErrCodePublicZoneVPCAssociation = "PublicZoneVPCAssociation"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The limit on the number of requests per second was exceeded.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeTooManyHealthChecks for service response error code
	// "TooManyHealthChecks".
	//
	// You have reached the maximum number of active health checks for an AWS account.
	// The default limit is 100. To request a higher limit, create a case (http://aws.amazon.com/route53-request)
	// with the AWS Support Center.
	ErrCodeTooManyHealthChecks = "TooManyHealthChecks"

	// ErrCodeTooManyHostedZones for service response error code
	// "TooManyHostedZones".
	//
	// This hosted zone can't be created because the hosted zone limit is exceeded.
	// To request a limit increase, go to the Amazon Route 53 Contact Us (http://aws.amazon.com/route53-request/)
	// page.
	ErrCodeTooManyHostedZones = "TooManyHostedZones"

	// ErrCodeTooManyTrafficPolicies for service response error code
	// "TooManyTrafficPolicies".
	//
	// You've created the maximum number of traffic policies that can be created
	// for the current AWS account. You can request an increase to the limit on
	// the Contact Us (http://aws.amazon.com/route53-request/) page.
	ErrCodeTooManyTrafficPolicies = "TooManyTrafficPolicies"

	// ErrCodeTooManyTrafficPolicyInstances for service response error code
	// "TooManyTrafficPolicyInstances".
	//
	// You've created the maximum number of traffic policy instances that can be
	// created for the current AWS account. You can request an increase to the limit
	// on the Contact Us (http://aws.amazon.com/route53-request/) page.
	ErrCodeTooManyTrafficPolicyInstances = "TooManyTrafficPolicyInstances"

	// ErrCodeTooManyVPCAssociationAuthorizations for service response error code
	// "TooManyVPCAssociationAuthorizations".
	//
	// You've created the maximum number of authorizations that can be created for
	// the specified hosted zone. To authorize another VPC to be associated with
	// the hosted zone, submit a DeleteVPCAssociationAuthorization request to remove
	// an existing authorization. To get a list of existing authorizations, submit
	// a ListVPCAssociationAuthorizations request.
	ErrCodeTooManyVPCAssociationAuthorizations = "TooManyVPCAssociationAuthorizations"

	// ErrCodeTrafficPolicyAlreadyExists for service response error code
	// "TrafficPolicyAlreadyExists".
	//
	// A traffic policy that has the same value for Name already exists.
	ErrCodeTrafficPolicyAlreadyExists = "TrafficPolicyAlreadyExists"

	// ErrCodeTrafficPolicyInUse for service response error code
	// "TrafficPolicyInUse".
	//
	// One or more traffic policy instances were created by using the specified
	// traffic policy.
	ErrCodeTrafficPolicyInUse = "TrafficPolicyInUse"

	// ErrCodeTrafficPolicyInstanceAlreadyExists for service response error code
	// "TrafficPolicyInstanceAlreadyExists".
	//
	// Traffic policy instance with given Id already exists.
	ErrCodeTrafficPolicyInstanceAlreadyExists = "TrafficPolicyInstanceAlreadyExists"

	// ErrCodeVPCAssociationAuthorizationNotFound for service response error code
	// "VPCAssociationAuthorizationNotFound".
	//
	// The VPC that you specified is not authorized to be associated with the hosted
	// zone.
	ErrCodeVPCAssociationAuthorizationNotFound = "VPCAssociationAuthorizationNotFound"

	// ErrCodeVPCAssociationNotFound for service response error code
	// "VPCAssociationNotFound".
	//
	// The specified VPC and hosted zone are not currently associated.
	ErrCodeVPCAssociationNotFound = "VPCAssociationNotFound"
)
