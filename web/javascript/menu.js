/* When the user clicks on the button, 
toggle between hiding and showing the dropdown content */
function showUserMenu() {
    document.getElementById("UserDropDown").classList.toggle("show");
  }

function showRoleMenu() {
    document.getElementById("RoleDropDown").classList.toggle("show");
  }
function showRoleBindingMenu() {
    document.getElementById("RoleBindingDropDown").classList.toggle("show");
}  
function showLogoutMenu() {
  document.getElementById("logoutdropdown").classList.toggle("show");
} 
  

  // Close the dropdown menu if the user clicks outside of it

  window.onclick = function(event) {
        var menu = ["User-DropDown", "Role-DropDown", "role-binding-dropdown","logout-dropdown"];
        for (i=0; i < 4; i++)
        {
            var dropdowns = document.getElementsByClassName(menu[i]);
            var j;
            for (j = 0; j < dropdowns.length; j++) {
                var openDropdown = dropdowns[j];
                if (openDropdown.classList.contains('show')) {
                openDropdown.classList.remove('show');
                }
            }
        }
}

window.addEventListener("pageshow", function(event) {
  var historyTraversal = event.persisted || (typeof window.performance !=
    "undefined" && window.performance.navigation.type === 2);
  if (historyTraversal) {
    // Handle page restore. 
    window.location.reload();
  }
});