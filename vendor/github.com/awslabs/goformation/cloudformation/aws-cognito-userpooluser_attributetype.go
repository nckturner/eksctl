package cloudformation

// AWSCognitoUserPoolUser_AttributeType AWS CloudFormation Resource (AWS::Cognito::UserPoolUser.AttributeType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html
type AWSCognitoUserPoolUser_AttributeType struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html#cfn-cognito-userpooluser-attributetype-name
	Name *StringIntrinsic `json:"Name,omitempty"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html#cfn-cognito-userpooluser-attributetype-value
	Value *StringIntrinsic `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolUser_AttributeType) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPoolUser.AttributeType"
}
