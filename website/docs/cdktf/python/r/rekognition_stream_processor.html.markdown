---
subcategory: "Rekognition"
layout: "aws"
page_title: "AWS: aws_rekognition_stream_processor"
description: |-
  Terraform resource for managing an AWS Rekognition Stream Processor.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_rekognition_stream_processor

Terraform resource for managing an AWS Rekognition Stream Processor.

~> This resource must be configured specifically for your use case, and not all options are compatible with one another. See [Stream Processor API documentation](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateStreamProcessor.html#rekognition-CreateStreamProcessor-request-Input) for configuration information.

~> Stream Processors configured for Face Recognition cannot have _any_ properties updated after the fact, and it will result in an AWS API error.

## Example Usage

### Label Detection

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.iam_role import IamRole
from imports.aws.kinesis_video_stream import KinesisVideoStream
from imports.aws.rekognition_stream_processor import RekognitionStreamProcessor
from imports.aws.s3_bucket import S3Bucket
from imports.aws.sns_topic import SnsTopic
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = KinesisVideoStream(self, "example",
            data_retention_in_hours=1,
            device_name="kinesis-video-device-name",
            media_type="video/h264",
            name="example-kinesis-input"
        )
        aws_s3_bucket_example = S3Bucket(self, "example_1",
            bucket="example-bucket"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_s3_bucket_example.override_logical_id("example")
        aws_sns_topic_example = SnsTopic(self, "example_2",
            name="example-topic"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_sns_topic_example.override_logical_id("example")
        aws_iam_role_example = IamRole(self, "example_3",
            assume_role_policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "rekognition.amazonaws.com"
                        }
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            inline_policy=[IamRoleInlinePolicy(
                name="Rekognition-Access",
                policy=Token.as_string(
                    Fn.jsonencode({
                        "Statement": [{
                            "Action": ["s3:PutObject"],
                            "Effect": "Allow",
                            "Resource": ["${" + aws_s3_bucket_example.arn + "}/*"]
                        }, {
                            "Action": ["sns:Publish"],
                            "Effect": "Allow",
                            "Resource": [aws_sns_topic_example.arn]
                        }, {
                            "Action": ["kinesis:Get*", "kinesis:DescribeStreamSummary"],
                            "Effect": "Allow",
                            "Resource": [example.arn]
                        }
                        ],
                        "Version": "2012-10-17"
                    }))
            )
            ],
            name="example-role"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_iam_role_example.override_logical_id("example")
        aws_rekognition_stream_processor_example = RekognitionStreamProcessor(self, "example_4",
            data_sharing_preference=[RekognitionStreamProcessorDataSharingPreference(
                opt_in=False
            )
            ],
            input=[RekognitionStreamProcessorInput(
                kinesis_video_stream=[RekognitionStreamProcessorInputKinesisVideoStream(
                    arn=example.arn
                )
                ]
            )
            ],
            name="example-processor",
            notification_channel=[RekognitionStreamProcessorNotificationChannel(
                sns_topic_arn=Token.as_string(aws_sns_topic_example.arn)
            )
            ],
            output=[RekognitionStreamProcessorOutput(
                s3_destination=[RekognitionStreamProcessorOutputS3Destination(
                    bucket=Token.as_string(aws_s3_bucket_example.bucket)
                )
                ]
            )
            ],
            role_arn=Token.as_string(aws_iam_role_example.arn),
            settings=[RekognitionStreamProcessorSettings(
                connected_home=[RekognitionStreamProcessorSettingsConnectedHome(
                    labels=["PERSON", "PET"]
                )
                ]
            )
            ]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_rekognition_stream_processor_example.override_logical_id("example")
```

### Face Detection Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.iam_role import IamRole
from imports.aws.kinesis_stream import KinesisStream
from imports.aws.kinesis_video_stream import KinesisVideoStream
from imports.aws.rekognition_collection import RekognitionCollection
from imports.aws.rekognition_stream_processor import RekognitionStreamProcessor
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = KinesisStream(self, "example",
            name="terraform-kinesis-example",
            shard_count=1
        )
        aws_kinesis_video_stream_example = KinesisVideoStream(self, "example_1",
            data_retention_in_hours=1,
            device_name="kinesis-video-device-name",
            media_type="video/h264",
            name="example-kinesis-input"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_kinesis_video_stream_example.override_logical_id("example")
        aws_rekognition_collection_example = RekognitionCollection(self, "example_2",
            collection_id="example-collection"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_rekognition_collection_example.override_logical_id("example")
        aws_iam_role_example = IamRole(self, "example_3",
            assume_role_policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "rekognition.amazonaws.com"
                        }
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            inline_policy=[IamRoleInlinePolicy(
                name="Rekognition-Access",
                policy=Token.as_string(
                    Fn.jsonencode({
                        "Statement": [{
                            "Action": ["kinesis:Get*", "kinesis:DescribeStreamSummary"],
                            "Effect": "Allow",
                            "Resource": [aws_kinesis_video_stream_example.arn]
                        }, {
                            "Action": ["kinesis:PutRecord"],
                            "Effect": "Allow",
                            "Resource": [example.arn]
                        }
                        ],
                        "Version": "2012-10-17"
                    }))
            )
            ],
            name="example-role"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_iam_role_example.override_logical_id("example")
        aws_rekognition_stream_processor_example = RekognitionStreamProcessor(self, "example_4",
            data_sharing_preference=[RekognitionStreamProcessorDataSharingPreference(
                opt_in=False
            )
            ],
            input=[RekognitionStreamProcessorInput(
                kinesis_video_stream=[RekognitionStreamProcessorInputKinesisVideoStream(
                    arn=Token.as_string(aws_kinesis_video_stream_example.arn)
                )
                ]
            )
            ],
            name="example-processor",
            output=[RekognitionStreamProcessorOutput(
                kinesis_data_stream=[RekognitionStreamProcessorOutputKinesisDataStream(
                    arn=example.arn
                )
                ]
            )
            ],
            regions_of_interest=[RekognitionStreamProcessorRegionsOfInterest(
                polygon=[RekognitionStreamProcessorRegionsOfInterestPolygon(
                    x=0.5,
                    y=0.5
                ), RekognitionStreamProcessorRegionsOfInterestPolygon(
                    x=0.5,
                    y=0.5
                ), RekognitionStreamProcessorRegionsOfInterestPolygon(
                    x=0.5,
                    y=0.5
                )
                ]
            )
            ],
            role_arn=Token.as_string(aws_iam_role_example.arn),
            settings=[RekognitionStreamProcessorSettings(
                face_search=[RekognitionStreamProcessorSettingsFaceSearch(
                    collection_id=Token.as_string(aws_rekognition_collection_example.id)
                )
                ]
            )
            ]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_rekognition_stream_processor_example.override_logical_id("example")
```

## Argument Reference

The following arguments are required:

* `input` - (Required) Input video stream. See [`input`](#input).
* `name` - (Required) The name of the Stream Processor.
* `output` - (Required) Kinesis data stream stream or Amazon S3 bucket location to which Amazon Rekognition Video puts the analysis results. See [`output`](#output).
* `role_arn` - (Required) The Amazon Resource Number (ARN) of the IAM role that allows access to the stream processor. The IAM role provides Rekognition read permissions for a Kinesis stream. It also provides write permissions to an Amazon S3 bucket and Amazon Simple Notification Service topic for a label detection stream processor. This is required for both face search and label detection stream processors.
* `settings` - (Required) Input parameters used in a streaming video analyzed by a stream processor. See [`settings`](#settings).

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `data_sharing_preference` - (Optional) See [`data_sharing_preference`](#data_sharing_preference).
* `kms_key_id` - (Optional) Optional parameter for label detection stream processors.
* `notification_channel` - (Optional) The Amazon Simple Notification Service topic to which Amazon Rekognition publishes the completion status. See [`notification_channel`](#notification_channel).
* `regions_of_interest` - (Optional) Specifies locations in the frames where Amazon Rekognition checks for objects or people. See [`regions_of_interest`](#regions_of_interest).
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### `input`

* `kinesis_video_stream` - (Optional) Kinesis input stream. See [`kinesis_video_stream`](#kinesis_video_stream).

### `kinesis_video_stream`

* `arn` - (Optional) ARN of the Kinesis video stream stream that streams the source video.

### `output`

* `kinesis_data_stream` - (Optional) The Amazon Kinesis Data Streams stream to which the Amazon Rekognition stream processor streams the analysis results. See [`kinesis_data_stream`](#kinesis_data_stream).
* `s3_destination` - (Optional) The Amazon S3 bucket location to which Amazon Rekognition publishes the detailed inference results of a video analysis operation. See [`s3_destination`](#s3_destination).

### `kinesis_data_stream`

* `arn` - (Optional) ARN of the output Amazon Kinesis Data Streams stream.

### `s3_destination`

* `bucket` - (Optional) Name of the Amazon S3 bucket you want to associate with the streaming video project.
* `key_prefixx` - (Optional) Prefix value of the location within the bucket that you want the information to be published to.

### `data_sharing_preference`

* `opt_in` - (Optional) Whether you are sharing data with Rekognition to improve model performance.

### `regions_of_interest`

* `bounding_box` - (Optional) Box representing a region of interest on screen. Only 1 per region is allowed. See [`bounding_box`](#bounding_box).
* `polygon` - (Optional) Shape made up of up to 10 Point objects to define a region of interest. See [`polygon`](#polygon).

### `bounding_box`

A region can only have a single `bounding_box`

* `height` - (Required) Height of the bounding box as a ratio of the overall image height.
* `wight` - (Required) Width of the bounding box as a ratio of the overall image width.
* `left` - (Required) Left coordinate of the bounding box as a ratio of overall image width.
* `top` - (Required) Top coordinate of the bounding box as a ratio of overall image height.

### `polygon`

If using `polygon`, a minimum of 3 per region is required, with a maximum of 10.

* `x` - (Required) The value of the X coordinate for a point on a Polygon.
* `y` - (Required) The value of the Y coordinate for a point on a Polygon.

### `notification_channel`

* `sns_topic_arn` - (Required) The Amazon Resource Number (ARN) of the Amazon Amazon Simple Notification Service topic to which Amazon Rekognition posts the completion status.

### `settings`

* `connected_home` - (Optional) Label detection settings to use on a streaming video. See [`connected_home`](#connected_home).
* `face_search` - (Optional) Input face recognition parameters for an Amazon Rekognition stream processor. See [`face_search`](#face_search).

### `connected_home`

* `labels` - (Required) Specifies what you want to detect in the video, such as people, packages, or pets. The current valid labels you can include in this list are: `PERSON`, `PET`, `PACKAGE`, and `ALL`.
* `min_confidence` - (Optional) Minimum confidence required to label an object in the video.

### `face_search`

* `collection_id` - (Optional) ID of a collection that contains faces that you want to search for.
* `face_match_threshold` - (Optional) Minimum face match confidence score that must be met to return a result for a recognized face.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Stream Processor.
* `stream_processor_arn` - (**Deprecated**) ARN of the Stream Processor.
  Use `arn` instead.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Rekognition Stream Processor using the `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.rekognition_stream_processor import RekognitionStreamProcessor
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        RekognitionStreamProcessor.generate_config_for_import(self, "example", "my-stream")
```

Using `terraform import`, import Rekognition Stream Processor using the `name`. For example:

```console
% terraform import aws_rekognition_stream_processor.example my-stream 
```

<!-- cache-key: cdktf-0.20.8 input-abe0591166b356690839bb3e9aa7a8f44be0b3d976ea807930a332854a221c26 -->