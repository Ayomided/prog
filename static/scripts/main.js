document.addEventListener('DOMContentLoaded', function () {
  // Trigger HTMX request to load default content
  const defaultOption = document.querySelector('.option[data-target="anyone"]');
  if (defaultOption) {
    defaultOption.click();
  }
});
