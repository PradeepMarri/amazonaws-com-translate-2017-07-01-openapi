package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CreateParallelDataResponse represents the CreateParallelDataResponse schema from the OpenAPI specification
type CreateParallelDataResponse struct {
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// TranslateTextRequest represents the TranslateTextRequest schema from the OpenAPI specification
type TranslateTextRequest struct {
	Text interface{} `json:"Text"`
	Settings interface{} `json:"Settings,omitempty"`
	Sourcelanguagecode interface{} `json:"SourceLanguageCode"`
	Targetlanguagecode interface{} `json:"TargetLanguageCode"`
	Terminologynames interface{} `json:"TerminologyNames,omitempty"`
}

// TerminologyProperties represents the TerminologyProperties schema from the OpenAPI specification
type TerminologyProperties struct {
	Directionality interface{} `json:"Directionality,omitempty"`
	Sourcelanguagecode interface{} `json:"SourceLanguageCode,omitempty"`
	Format interface{} `json:"Format,omitempty"`
	Message interface{} `json:"Message,omitempty"`
	Skippedtermcount interface{} `json:"SkippedTermCount,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Encryptionkey interface{} `json:"EncryptionKey,omitempty"`
	Lastupdatedat interface{} `json:"LastUpdatedAt,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Targetlanguagecodes interface{} `json:"TargetLanguageCodes,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Sizebytes interface{} `json:"SizeBytes,omitempty"`
	Termcount interface{} `json:"TermCount,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
	Languagecode interface{} `json:"LanguageCode"`
	Languagename interface{} `json:"LanguageName"`
}

// DescribeTextTranslationJobResponse represents the DescribeTextTranslationJobResponse schema from the OpenAPI specification
type DescribeTextTranslationJobResponse struct {
	Texttranslationjobproperties interface{} `json:"TextTranslationJobProperties,omitempty"`
}

// JobDetails represents the JobDetails schema from the OpenAPI specification
type JobDetails struct {
	Documentswitherrorscount interface{} `json:"DocumentsWithErrorsCount,omitempty"`
	Inputdocumentscount interface{} `json:"InputDocumentsCount,omitempty"`
	Translateddocumentscount interface{} `json:"TranslatedDocumentsCount,omitempty"`
}

// ImportTerminologyResponse represents the ImportTerminologyResponse schema from the OpenAPI specification
type ImportTerminologyResponse struct {
	Terminologyproperties interface{} `json:"TerminologyProperties,omitempty"`
	Auxiliarydatalocation interface{} `json:"AuxiliaryDataLocation,omitempty"`
}

// TerminologyDataLocation represents the TerminologyDataLocation schema from the OpenAPI specification
type TerminologyDataLocation struct {
	Location interface{} `json:"Location"`
	Repositorytype interface{} `json:"RepositoryType"`
}

// ListParallelDataRequest represents the ListParallelDataRequest schema from the OpenAPI specification
type ListParallelDataRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tags interface{} `json:"Tags"`
}

// ListTextTranslationJobsRequest represents the ListTextTranslationJobsRequest schema from the OpenAPI specification
type ListTextTranslationJobsRequest struct {
	Filter interface{} `json:"Filter,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// StartTextTranslationJobResponse represents the StartTextTranslationJobResponse schema from the OpenAPI specification
type StartTextTranslationJobResponse struct {
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
}

// GetTerminologyRequest represents the GetTerminologyRequest schema from the OpenAPI specification
type GetTerminologyRequest struct {
	Name interface{} `json:"Name"`
	Terminologydataformat interface{} `json:"TerminologyDataFormat,omitempty"`
}

// UpdateParallelDataResponse represents the UpdateParallelDataResponse schema from the OpenAPI specification
type UpdateParallelDataResponse struct {
	Latestupdateattemptstatus interface{} `json:"LatestUpdateAttemptStatus,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Latestupdateattemptat interface{} `json:"LatestUpdateAttemptAt,omitempty"`
}

// TextTranslationJobFilter represents the TextTranslationJobFilter schema from the OpenAPI specification
type TextTranslationJobFilter struct {
	Jobname interface{} `json:"JobName,omitempty"`
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Submittedaftertime interface{} `json:"SubmittedAfterTime,omitempty"`
	Submittedbeforetime interface{} `json:"SubmittedBeforeTime,omitempty"`
}

// ImportTerminologyRequest represents the ImportTerminologyRequest schema from the OpenAPI specification
type ImportTerminologyRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Encryptionkey interface{} `json:"EncryptionKey,omitempty"`
	Mergestrategy interface{} `json:"MergeStrategy"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Terminologydata interface{} `json:"TerminologyData"`
}

// CreateParallelDataRequest represents the CreateParallelDataRequest schema from the OpenAPI specification
type CreateParallelDataRequest struct {
	Paralleldataconfig interface{} `json:"ParallelDataConfig"`
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken"`
	Description interface{} `json:"Description,omitempty"`
	Encryptionkey EncryptionKey `json:"EncryptionKey,omitempty"` // The encryption key used to encrypt this object.
	Name interface{} `json:"Name"`
}

// InputDataConfig represents the InputDataConfig schema from the OpenAPI specification
type InputDataConfig struct {
	Contenttype interface{} `json:"ContentType"`
	S3uri interface{} `json:"S3Uri"`
}

// GetParallelDataRequest represents the GetParallelDataRequest schema from the OpenAPI specification
type GetParallelDataRequest struct {
	Name interface{} `json:"Name"`
}

// TranslateDocumentResponse represents the TranslateDocumentResponse schema from the OpenAPI specification
type TranslateDocumentResponse struct {
	Translateddocument interface{} `json:"TranslatedDocument"`
	Appliedsettings TranslationSettings `json:"AppliedSettings,omitempty"` // Settings to configure your translation output, including the option to set the formality level of the output text and the option to mask profane words and phrases.
	Appliedterminologies interface{} `json:"AppliedTerminologies,omitempty"`
	Sourcelanguagecode interface{} `json:"SourceLanguageCode"`
	Targetlanguagecode interface{} `json:"TargetLanguageCode"`
}

// AppliedTerminology represents the AppliedTerminology schema from the OpenAPI specification
type AppliedTerminology struct {
	Terms interface{} `json:"Terms,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// StopTextTranslationJobRequest represents the StopTextTranslationJobRequest schema from the OpenAPI specification
type StopTextTranslationJobRequest struct {
	Jobid interface{} `json:"JobId"`
}

// StartTextTranslationJobRequest represents the StartTextTranslationJobRequest schema from the OpenAPI specification
type StartTextTranslationJobRequest struct {
	Clienttoken interface{} `json:"ClientToken"`
	Inputdataconfig interface{} `json:"InputDataConfig"`
	Jobname interface{} `json:"JobName,omitempty"`
	Terminologynames interface{} `json:"TerminologyNames,omitempty"`
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn"`
	Outputdataconfig interface{} `json:"OutputDataConfig"`
	Settings interface{} `json:"Settings,omitempty"`
	Paralleldatanames interface{} `json:"ParallelDataNames,omitempty"`
	Sourcelanguagecode interface{} `json:"SourceLanguageCode"`
	Targetlanguagecodes interface{} `json:"TargetLanguageCodes"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// DescribeTextTranslationJobRequest represents the DescribeTextTranslationJobRequest schema from the OpenAPI specification
type DescribeTextTranslationJobRequest struct {
	Jobid interface{} `json:"JobId"`
}

// UpdateParallelDataRequest represents the UpdateParallelDataRequest schema from the OpenAPI specification
type UpdateParallelDataRequest struct {
	Name interface{} `json:"Name"`
	Paralleldataconfig interface{} `json:"ParallelDataConfig"`
	Clienttoken interface{} `json:"ClientToken"`
	Description interface{} `json:"Description,omitempty"`
}

// OutputDataConfig represents the OutputDataConfig schema from the OpenAPI specification
type OutputDataConfig struct {
	Encryptionkey EncryptionKey `json:"EncryptionKey,omitempty"` // The encryption key used to encrypt this object.
	S3uri interface{} `json:"S3Uri"`
}

// ParallelDataConfig represents the ParallelDataConfig schema from the OpenAPI specification
type ParallelDataConfig struct {
	Format interface{} `json:"Format"`
	S3uri interface{} `json:"S3Uri"`
}

// ListLanguagesRequest represents the ListLanguagesRequest schema from the OpenAPI specification
type ListLanguagesRequest struct {
	Displaylanguagecode interface{} `json:"DisplayLanguageCode,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// EncryptionKey represents the EncryptionKey schema from the OpenAPI specification
type EncryptionKey struct {
	Id interface{} `json:"Id"`
	TypeField interface{} `json:"Type"`
}

// TranslationSettings represents the TranslationSettings schema from the OpenAPI specification
type TranslationSettings struct {
	Formality interface{} `json:"Formality,omitempty"`
	Profanity interface{} `json:"Profanity,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// Term represents the Term schema from the OpenAPI specification
type Term struct {
	Targettext interface{} `json:"TargetText,omitempty"`
	Sourcetext interface{} `json:"SourceText,omitempty"`
}

// TranslateDocumentRequest represents the TranslateDocumentRequest schema from the OpenAPI specification
type TranslateDocumentRequest struct {
	Targetlanguagecode interface{} `json:"TargetLanguageCode"`
	Terminologynames interface{} `json:"TerminologyNames,omitempty"`
	Document interface{} `json:"Document"`
	Settings TranslationSettings `json:"Settings,omitempty"` // Settings to configure your translation output, including the option to set the formality level of the output text and the option to mask profane words and phrases.
	Sourcelanguagecode interface{} `json:"SourceLanguageCode"`
}

// TranslatedDocument represents the TranslatedDocument schema from the OpenAPI specification
type TranslatedDocument struct {
	Content interface{} `json:"Content"`
}

// ParallelDataDataLocation represents the ParallelDataDataLocation schema from the OpenAPI specification
type ParallelDataDataLocation struct {
	Repositorytype interface{} `json:"RepositoryType"`
	Location interface{} `json:"Location"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// ListParallelDataResponse represents the ListParallelDataResponse schema from the OpenAPI specification
type ListParallelDataResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Paralleldatapropertieslist interface{} `json:"ParallelDataPropertiesList,omitempty"`
}

// ListLanguagesResponse represents the ListLanguagesResponse schema from the OpenAPI specification
type ListLanguagesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Displaylanguagecode interface{} `json:"DisplayLanguageCode,omitempty"`
	Languages interface{} `json:"Languages,omitempty"`
}

// TranslateTextResponse represents the TranslateTextResponse schema from the OpenAPI specification
type TranslateTextResponse struct {
	Targetlanguagecode interface{} `json:"TargetLanguageCode"`
	Translatedtext interface{} `json:"TranslatedText"`
	Appliedsettings interface{} `json:"AppliedSettings,omitempty"`
	Appliedterminologies interface{} `json:"AppliedTerminologies,omitempty"`
	Sourcelanguagecode interface{} `json:"SourceLanguageCode"`
}

// TextTranslationJobProperties represents the TextTranslationJobProperties schema from the OpenAPI specification
type TextTranslationJobProperties struct {
	Outputdataconfig interface{} `json:"OutputDataConfig,omitempty"`
	Targetlanguagecodes interface{} `json:"TargetLanguageCodes,omitempty"`
	Message interface{} `json:"Message,omitempty"`
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig,omitempty"`
	Submittedtime interface{} `json:"SubmittedTime,omitempty"`
	Jobdetails interface{} `json:"JobDetails,omitempty"`
	Terminologynames interface{} `json:"TerminologyNames,omitempty"`
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Settings interface{} `json:"Settings,omitempty"`
	Sourcelanguagecode interface{} `json:"SourceLanguageCode,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Paralleldatanames interface{} `json:"ParallelDataNames,omitempty"`
	Jobname interface{} `json:"JobName,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// ParallelDataProperties represents the ParallelDataProperties schema from the OpenAPI specification
type ParallelDataProperties struct {
	Sourcelanguagecode interface{} `json:"SourceLanguageCode,omitempty"`
	Targetlanguagecodes interface{} `json:"TargetLanguageCodes,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Lastupdatedat interface{} `json:"LastUpdatedAt,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Importedrecordcount interface{} `json:"ImportedRecordCount,omitempty"`
	Message interface{} `json:"Message,omitempty"`
	Failedrecordcount interface{} `json:"FailedRecordCount,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Importeddatasize interface{} `json:"ImportedDataSize,omitempty"`
	Encryptionkey EncryptionKey `json:"EncryptionKey,omitempty"` // The encryption key used to encrypt this object.
	Name interface{} `json:"Name,omitempty"`
	Paralleldataconfig interface{} `json:"ParallelDataConfig,omitempty"`
	Latestupdateattemptat interface{} `json:"LatestUpdateAttemptAt,omitempty"`
	Latestupdateattemptstatus interface{} `json:"LatestUpdateAttemptStatus,omitempty"`
	Skippedrecordcount interface{} `json:"SkippedRecordCount,omitempty"`
}

// DeleteParallelDataRequest represents the DeleteParallelDataRequest schema from the OpenAPI specification
type DeleteParallelDataRequest struct {
	Name interface{} `json:"Name"`
}

// ListTerminologiesResponse represents the ListTerminologiesResponse schema from the OpenAPI specification
type ListTerminologiesResponse struct {
	Terminologypropertieslist interface{} `json:"TerminologyPropertiesList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteTerminologyRequest represents the DeleteTerminologyRequest schema from the OpenAPI specification
type DeleteTerminologyRequest struct {
	Name interface{} `json:"Name"`
}

// GetParallelDataResponse represents the GetParallelDataResponse schema from the OpenAPI specification
type GetParallelDataResponse struct {
	Latestupdateattemptauxiliarydatalocation interface{} `json:"LatestUpdateAttemptAuxiliaryDataLocation,omitempty"`
	Paralleldataproperties interface{} `json:"ParallelDataProperties,omitempty"`
	Auxiliarydatalocation interface{} `json:"AuxiliaryDataLocation,omitempty"`
	Datalocation interface{} `json:"DataLocation,omitempty"`
}

// Document represents the Document schema from the OpenAPI specification
type Document struct {
	Content interface{} `json:"Content"`
	Contenttype interface{} `json:"ContentType"`
}

// ListTerminologiesRequest represents the ListTerminologiesRequest schema from the OpenAPI specification
type ListTerminologiesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// StopTextTranslationJobResponse represents the StopTextTranslationJobResponse schema from the OpenAPI specification
type StopTextTranslationJobResponse struct {
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
}

// DeleteParallelDataResponse represents the DeleteParallelDataResponse schema from the OpenAPI specification
type DeleteParallelDataResponse struct {
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tagkeys interface{} `json:"TagKeys"`
}

// TerminologyData represents the TerminologyData schema from the OpenAPI specification
type TerminologyData struct {
	Directionality interface{} `json:"Directionality,omitempty"`
	File interface{} `json:"File"`
	Format interface{} `json:"Format"`
}

// GetTerminologyResponse represents the GetTerminologyResponse schema from the OpenAPI specification
type GetTerminologyResponse struct {
	Terminologyproperties interface{} `json:"TerminologyProperties,omitempty"`
	Auxiliarydatalocation interface{} `json:"AuxiliaryDataLocation,omitempty"`
	Terminologydatalocation interface{} `json:"TerminologyDataLocation,omitempty"`
}

// ListTextTranslationJobsResponse represents the ListTextTranslationJobsResponse schema from the OpenAPI specification
type ListTextTranslationJobsResponse struct {
	Texttranslationjobpropertieslist interface{} `json:"TextTranslationJobPropertiesList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}
