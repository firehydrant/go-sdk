# ListIncidentEventsResponse


## Fields

| Field                                                                                                                                 | Type                                                                                                                                  | Required                                                                                                                              | Description                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                                                                            | [components.HTTPMetadata](../../models/components/httpmetadata.md)                                                                    | :heavy_check_mark:                                                                                                                    | N/A                                                                                                                                   |
| `IncidentEventEntityPaginated`                                                                                                        | [*components.IncidentEventEntityPaginated](../../models/components/incidentevententitypaginated.md)                                   | :heavy_minus_sign:                                                                                                                    | List all events for an incident. An event is a timeline entry. This can be filtered with params to retrieve events of a certain type. |