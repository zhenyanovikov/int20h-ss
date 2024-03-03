import { useState } from "react";
import { useDebounce } from "@uidotdev/usehooks";
import { useGetStudents } from "../../../api/student";
import StudentsTemplate from "../../templates/StudentsTemplate/StudentsTemplate";

function StudentsPage() {
  const [name, setName] = useState("");
  const [groupId, setGroupId] = useState("");

  const debouncedName = useDebounce(name, 300);

  let params = {};
  if (debouncedName) {
    params.name = debouncedName;
  }
  if (groupId) {
    params.groupId = groupId;
  }

  const { data: studentsData, isLoading: isGetStudentsLoading } =
    useGetStudents(params);

  return (
    <StudentsTemplate
      studentsData={studentsData}
      isGetStudentsLoading={isGetStudentsLoading}
      name={name}
      onChangeName={setName}
      groupId={groupId}
      onChangeGroupId={setGroupId}
    />
  );
}

export default StudentsPage;
