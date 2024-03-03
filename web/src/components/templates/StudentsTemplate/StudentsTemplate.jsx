import { Stack, CircularProgress } from "@mui/material";
import Students from "../../organisms/Students/Students";
import InviteStudentForm from "../../organisms/InviteStudentForm/InviteStudentForm";
import FilterStudents from "../../organisms/FilterStudents/FilterStudents";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";
import Empty from "../../molecules/Empty/Empty";

function StudentsTemplate({
  studentsData,
  isGetStudentsLoading,
  name,
  onChangeName,
  groupId,
  onChangeGroupId,
}) {
  return (
    <AdminTemplate>
      <Stack spacing={8}>
        <InviteStudentForm />

        <FilterStudents
          name={name}
          onChangeName={onChangeName}
          groupId={groupId}
          onChangeGroupId={onChangeGroupId}
        />

        {renderContent()}
      </Stack>
    </AdminTemplate>
  );

  function renderContent() {
    if (isGetStudentsLoading || !studentsData) {
      return (
        <Stack alignItems="center">
          <CircularProgress />
        </Stack>
      );
    }

    if (!studentsData.length) {
      return <Empty />;
    }

    return <Students students={studentsData} />;
  }
}

export default StudentsTemplate;
