resource "google_app_engine_application" "app" {
  project       = var.project_id
  location_id   = var.location_id
  database_type = "CLOUD_FIRESTORE"
}

resource "google_firestore_document" "articles" {
  project     = var.project_id
  document_id = "article_1"
  collection  = "articles"
  #fields = "{\"something\":{\"mapValue\":{\"fields\":{\"akey\":{\"stringValue\":\"avalue\"}}}}}"
  fields = "{\"title\":{\"stringValue\":\"atitle\"},\"subtitle\":{\"stringValue\":\"asubtitle\"},\"blocks\":{\"mapValue\":{\"fields\":{\"type\":{\"stringValue\":\"text\"},\"value\":{\"stringValue\": \"article paragraph\"},\"index\":{\"integerValue\":\"0\"}}}},\"banner\":{\"stringValue\":\"url\"},\"tags\":{\"arrayValue\":{\"values\":[{\"stringValue\":\"atag\"}]}},\"creationDate\":{\"timestampValue\":\"2021-12-12T22:57:03Z\"},\"lastChangeDate\":{\"timestampValue\":\"2021-12-12T22:57:03Z\"},\"authorID\":{\"stringValue\":\"authorID\"}}"
}
