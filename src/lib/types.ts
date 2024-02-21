/**
 *
 * @param {BackendObject} Object Das Objekt, welches auch so im Backend gespeichert wird
 * @param {number} object.ID Die ID des Objekts
 * @param {string} object.Name Name des Objekts
 * @param {string} object.Type Type des Objekts
 * @param {string} object.Data Daten des Objekts
 * @param {number} object.FolderID ID des Ordners, in dem das Objekt liegt
 *
 */
interface BackendObject {
    ID: number;
    Name: string;
    Type: string;
    Data: string;
    FolderID: number;
}


interface BackendFolder{
    ID: number;
    Name: string;
    ParentID: number;
    Children: BackendFolder[];
    Objects: BackendObject[];
}

export { BackendObject, BackendFolder };