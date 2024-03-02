import { Stack } from "@mui/material";
import Faculty from "../../molecules/Faculty/Faculty";

function Faculties({ faculties }) {
  return (
    <Stack spacing={2}>
      {faculties.map((faculty) => (
        <Faculty key={faculty.id} faculty={faculty} />
      ))}
    </Stack>
  );
}

export default Faculties;
