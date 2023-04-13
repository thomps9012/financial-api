export default function Footer() {
  return (
    <footer>
      <p>© {new Date().getFullYear()}</p>
      <div
        style={{ display: "flex", flexDirection: "column", textAlign: "right" }}
      >
        <a href="https://tszlau.com/">
          <p>
            Designed by
            <br />
            tszlau WebDesign
          </p>
        </a>
      </div>
    </footer>
  );
}
