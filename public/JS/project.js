
let projects = []

function addProject(event) {
    event.preventDefault()
    
    let title = document.getElementById('input-project').value
    let getStartDate = document.getElementById('input-start').value
    let getEndDate = document.getElementById('input-end').value
    let desc = document.getElementById('input-desc').value
    let tech1 = document.getElementById('node').checked
    let tech2 = document.getElementById('react').checked
    let tech3 = document.getElementById('next').checked
    let tech4 = document.getElementById('type').checked
    let image = document.getElementById('input-img').files
    
    image = URL.createObjectURL(image[0])

    getEndDate = new Date(getEndDate)
    getStartDate = new Date(getStartDate)


    let project = {
        title,
        getStartDate,
        getEndDate,
        desc,
        tech1,
        tech2,
        tech3,
        tech4,
        image
    }

    projects.push(project)

    renderProject()
}

function renderProject() {
    
    document.getElementById('list-myproject').innerHTML = ''

    for (let i = 0; i < projects.length; i++) {
        
        document.getElementById('list-myproject').innerHTML += `

        <div class="project-list-item">
                <a href="project-detail.html">

                    <div class="card-img">
                        <img src="${projects[i].image}">
                    </div>

                    <div class="card-title">
                        <h3>${projects[i].title}</h3>
                    </div>

                    <div class="card-drt">
                        <p>${monthDuration(projects[i].getEndDate, projects[i].getStartDate)}</p>
                    </div>

                    <div class="card-desc">
                        <p>${projects[i].desc}</p>
                    </div>

                    <div class="card-icon">
                        ${projects[i].tech1?`<img src="assets/IMG/nodejs.png"/>` : ""}
                        ${projects[i].tech2?`<img src="assets/IMG/react.png"/>` : ""}
                        ${projects[i].tech3?`<img src="assets/IMG/nextjs.png"/>` : ""}
                        ${projects[i].tech4?`<img src="assets/IMG/typescript.png"/>` : ""}
                        
                    </div>

                </a>

                <div class="card-btn">
                    <div class="edit-btn">
                        <button>edit</button>
                    </div>
                    <div class="del-btn">
                        <button>delete</button>
                    </div>
                </div>

            </div>
        
        `
    
    }

}

function monthDuration(endDate, startDate) {
    // month
    // year
    // monthDistance

    let endMonth = endDate.getMonth()
    let startMonth = startDate.getMonth()
    let endYear = endDate.getFullYear()
    let startYear = startDate.getFullYear()

    if(startYear == endYear) {
        if(startMonth == endMonth){
            month = 1
            return 'durasi ' + month + ' bulan'
        }else{
            month = endMonth - startMonth
            return 'durasi ' + month + ' bulan'
        }
    } 
  
    
    if (endYear > startYear) {
        if (endYear - startYear == 1) {
            if (startMonth == endMonth) {
                return 'durasi 1 tahun'
            }else if (startMonth > endMonth) {
                month = (startMonth - endMonth - 12) * -1
                return 'durasi ' + month + ' bulan'
            }else if (startMonth < endMonth) {
                month = endMonth - startMonth
                return 'durasi 1 tahun ' + month + ' bulan'
            }
        }else{
            year = endYear - startYear
            if (startMonth == endMonth) {
                return 'durasi ' + year + ' tahun '
            }else if (startMonth > endMonth) {
                year -= 1
                month = (startMonth - endMonth - 12) * -1
                return 'durasi ' + year + ' tahun ' + month + ' bulan'
            }else if (startMonth < endMonth) {
                month = endMonth - startMonth
                return 'durasi ' + year + ' tahun ' + month + ' bulan'
            }
        }
    }

}


