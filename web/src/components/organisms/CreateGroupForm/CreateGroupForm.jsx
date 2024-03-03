import {
  Stack,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  TextField,
  Button,
  Box,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { useFormik } from "formik";
import { useGetFaculties } from "../../../api/faculty";
import { useCreateGroup } from "../../../api/group";
import { getCreateGroupValidationSchema } from "../../../helpers/validation";

function CreateGroupForm() {
  const { t } = useTranslation();
  const { data: facultiesData } = useGetFaculties();
  const { mutate: createGroupMutate, isLoading: isCreateGroupLoading } =
    useCreateGroup();

  const formik = useFormik({
    initialValues: {
      facultyId: "",
      yearStart: "",
      name: "",
    },
    validationSchema: getCreateGroupValidationSchema(),
    onSubmit: handleSubmit,
  });

  const faculties = facultiesData || [];

  return (
    <form onSubmit={formik.handleSubmit}>
      <Stack
        sx={{
          flexDirection: {
            sx: "column",
            md: "row",
          },
          gap: 2,
        }}
      >
        <Box sx={{ width: { sx: "100%", md: "184px" } }}>
          <FormControl size="small" fullWidth>
            <InputLabel id="facultyId">
              {t("organisms.createGroupForm.form.facultyId.label")}
            </InputLabel>
            <Select
              labelId="facultyId"
              id="facultyId"
              name="facultyId"
              value={formik.values.facultyId}
              label={t("organisms.createGroupForm.form.facultyId.label")}
              onChange={formik.handleChange}
            >
              {faculties.map((faculty) => (
                <MenuItem key={faculty.id} value={faculty.id}>
                  {faculty.name}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Box>
        <TextField
          id="name"
          name="name"
          size="small"
          label={t("organisms.createGroupForm.form.name.label")}
          value={formik.values.name}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.name && !!formik.errors.name}
          helperText={formik.touched.name && t(formik.errors.name)}
        />
        <TextField
          id="yearStart"
          name="yearStart"
          size="small"
          type="number"
          label={t("organisms.createGroupForm.form.yearStart.label")}
          value={formik.values.yearStart}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.yearStart && !!formik.errors.yearStart}
          helperText={formik.touched.yearStart && t(formik.errors.yearStart)}
        />
        <Stack>
          <Button
            variant="contained"
            type="submit"
            disabled={isCreateGroupLoading}
          >
            {t("organisms.createGroupForm.form.button.submit")}
          </Button>
        </Stack>
      </Stack>
    </form>
  );

  function handleSubmit(values, actions) {
    createGroupMutate(values);

    actions.resetForm();
  }
}

export default CreateGroupForm;
