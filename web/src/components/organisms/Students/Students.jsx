import { Stack } from "@mui/material";
import Student from "../../molecules/Student/Student";

function Students({ students }) {
  return (
    <Stack spacing={2}>
      {students.map((student) => (
        <Student key={student.id} student={student} />
      ))}
    </Stack>
  );
}

export default Students;
