# ResolveIncidentRequestBody


## Fields

| Field                                                                                                                                                                                                        | Type                                                                                                                                                                                                         | Required                                                                                                                                                                                                     | Description                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `Milestone`                                                                                                                                                                                                  | **string*                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                           | The slug of any milestone in the post-incident or closed phase to set on the incident (and its children, if `resolve_children` os set). Must be one of the configured milestones available on this incident. |