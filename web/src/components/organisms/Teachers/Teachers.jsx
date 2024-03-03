import { Stack } from "@mui/material";
import Teacher from "../../molecules/Teacher/Teacher";

function Teachers({ teachers }) {
  return (
    <Stack spacing={2}>
      {teachers.map((teacher) => (
        <Teacher key={teacher.id} teacher={teacher} />
      ))}
    </Stack>
  );
}

export default Teachers;
